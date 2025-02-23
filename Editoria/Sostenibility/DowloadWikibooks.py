import requests
from bs4 import BeautifulSoup

# Per scaricare il contenuto di un libro da Wikibooks moficare l'URL e il nome del file di output

def download_wikibooks_content(url, output_file):
    response = requests.get(url)
    if response.status_code != 200:
        print(f"Errore nel caricamento della pagina: {response.status_code}")
        return

    soup = BeautifulSoup(response.text, "html.parser")
    content = ""

    # Estrarre il titolo e il contenuto
    title = soup.find("h1").text
    content += f"# {title}\n\n"
    paragraphs = soup.find_all("p")

    for paragraph in paragraphs:
        content += paragraph.text + "\n\n"

    # Salva il contenuto in un file Markdown
    with open(output_file, "w", encoding="utf-8") as file:
        file.write(content)
    print(f"Contenuto salvato in {output_file}")

if __name__ == "__main__":
    url = "https://it.wikibooks.org/wiki/Tecnologia/Regole_per_contribuire"
    output_file = "./docs/Regole_Ambientali.md"
    download_wikibooks_content(url, output_file)
