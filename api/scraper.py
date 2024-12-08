import requests
from bs4 import BeautifulSoup


def extract(url):
    res = requests.get(url)
    if res.status_code != 200:
        return None

    soup = BeautifulSoup(res.content, "html.parser")

    title = soup.find("span", class_="mw-page-title-main")
    content = soup.find_all("p")

    text = ""
    if title is not None:
        text += title.get_text(strip=True).replace("\n", " ").replace("\r", " ")

    for para in content:
        if len(text.split()) > 600:
            return text
        if para is not None:
            text += para.get_text()

    return text
