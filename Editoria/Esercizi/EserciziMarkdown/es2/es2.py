def utf8_to_text(utf8_codes):
    return bytes(utf8_codes).decode('utf-8')

#Test
utf8_codes = [0xc2, 0xab]
text = utf8_to_text(utf8_codes)
print(f"Decoded text: {text}")