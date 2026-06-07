import logging
import requests
from bs4 import BeautifulSoup

logger = logging.getLogger(__name__)

# Realistic browser headers to avoid being blocked
HEADERS = {
    "User-Agent": (
        "Mozilla/5.0 (Windows NT 10.0; Win64; x64) "
        "AppleWebKit/537.36 (KHTML, like Gecko) "
        "Chrome/125.0.0.0 Safari/537.36"
    ),
    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
    "Accept-Language": "en-US,en;q=0.5",
    "Accept-Encoding": "gzip, deflate",
    "DNT": "1",
    "Connection": "keep-alive",
    "Upgrade-Insecure-Requests": "1",
}

REQUEST_TIMEOUT = 15  # seconds


def extract(url: str) -> str | None:
    """
    Fetch and extract readable text content from a URL.

    Safely downloads the page with realistic browser headers, parses
    the HTML, and extracts text from the most relevant content
    containers. Returns up to ~600 words of clean text, or None on
    failure.
    """
    logger.info("Fetching URL: %s", url)

    try:
        res = requests.get(url, headers=HEADERS, timeout=REQUEST_TIMEOUT)
        res.raise_for_status()
    except requests.exceptions.Timeout:
        logger.error("Request timed out for URL: %s", url)
        return None
    except requests.exceptions.RequestException as e:
        logger.error("Request failed for URL %s: %s", url, e)
        return None

    logger.debug("HTTP %d for %s (%d bytes)", res.status_code, url, len(res.content))

    soup = BeautifulSoup(res.content, "html.parser")

    # Remove non-content elements
    for tag in soup(["script", "style", "nav", "footer", "header", "aside", "noscript"]):
        tag.decompose()

    # Try the most semantic containers first
    text = ""
    for selector in ["article", "main", "[role='main']", "body"]:
        container = soup.select_one(selector)
        if container:
            raw = container.get_text(separator="\n", strip=True)
            if len(raw.split()) > 50:
                text = raw
                logger.debug("Extracted content via <%s> selector", selector)
                break

    # Fallback: if nothing matched, pull all <p> text
    if not text:
        paragraphs = soup.find_all("p")
        text = "\n".join(p.get_text(strip=True) for p in paragraphs if p.get_text(strip=True))
        logger.debug("Extracted content via <p> fallback")

    # Truncate to roughly 600 words (keep it reasonable for summarisation)
    words = text.split()
    if len(words) > 600:
        text = " ".join(words[:600])
        logger.debug("Truncated to 600 words")

    word_count = len(text.split())
    logger.info("Extracted %d words from %s", word_count, url)
    return text if text.strip() else None
