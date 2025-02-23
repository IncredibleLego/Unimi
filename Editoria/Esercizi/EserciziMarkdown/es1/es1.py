def text_to_ascii(text):
    return [ord(char) for char in text]

#test
text = "Buongiorno"
ascii_codes = text_to_ascii(text)
print(f"ASCII codes: {ascii_codes}")