import yaml
import json
from lxml import etree

# Funzione per leggere il file YAML
def read_yaml(yaml_file):
    with open(yaml_file, 'r') as file:
        return yaml.safe_load(file)

# Funzione per convertire YAML a Schema.org
def convert_to_schema_org(data):
    schema_org = {
        "@context": "http://schema.org",
        "@type": "Book",
        "name": data.get("title"),
        "author": {
            "@type": "Person",
            "name": data.get("author")
        },
        "datePublished": data.get("publication_date"),
        "isbn": data.get("isbn"),
        "publisher": {
            "@type": "Organization",
            "name": data.get("publisher")
        },
    }
    return schema_org

# Funzione per convertire YAML a ONIX (semplificato)
def convert_to_onix(data):
    root = etree.Element("Product")
    record_reference = etree.SubElement(root, "RecordReference")
    record_reference.text = data.get("isbn")

    notification_type = etree.SubElement(root, "NotificationType")
    notification_type.text = "03"  # esempio: 03=Nuova edizione

    product_identifier = etree.SubElement(root, "ProductIdentifier")
    product_id_type = etree.SubElement(product_identifier, "ProductIDType")
    product_id_type.text = "15"  # ISBN-13
    id_value = etree.SubElement(product_identifier, "IDValue")
    id_value.text = data.get("isbn")

    descriptive_detail = etree.SubElement(root, "DescriptiveDetail")
    title_detail = etree.SubElement(descriptive_detail, "TitleDetail")
    title_type = etree.SubElement(title_detail, "TitleType")
    title_type.text = "01"  # Titolo principale
    title_element = etree.SubElement(title_detail, "TitleElement")
    title_text = etree.SubElement(title_element, "TitleText")
    title_text.text = data.get("title")

    contributor = etree.SubElement(descriptive_detail, "Contributor")
    contributor_role = etree.SubElement(contributor, "ContributorRole")
    contributor_role.text = "A01"  # Autore
    person_name = etree.SubElement(contributor, "PersonName")
    person_name.text = data.get("author")

    publishing_detail = etree.SubElement(root, "PublishingDetail")
    publisher = etree.SubElement(publishing_detail, "Publisher")
    publisher_name = etree.SubElement(publisher, "PublisherName")
    publisher_name.text = data.get("publisher")

    publishing_date = etree.SubElement(publishing_detail, "PublishingDate")
    publishing_date_role = etree.SubElement(publishing_date, "PublishingDateRole")
    publishing_date_role.text = "01"  # Data di pubblicazione
    date = etree.SubElement(publishing_date, "Date")
    date.text = data.get("publication_date")

    return etree.tostring(root, pretty_print=True, xml_declaration=True, encoding="UTF-8")

# Funzione principale per eseguire la conversione
def main(yaml_file):
    data = read_yaml(yaml_file)

    # Convertire a Schema.org
    schema_org = convert_to_schema_org(data)
    with open('output_schema_org.json', 'w') as file:
        json.dump(schema_org, file, indent=2)

    # Convertire a ONIX
    onix = convert_to_onix(data)
    with open('output_onix.xml', 'wb') as file:
        file.write(onix)

if __name__ == "__main__":
    yaml_file = 'Napoleone_Bonaparte_metadata.yaml'
    main(yaml_file)
