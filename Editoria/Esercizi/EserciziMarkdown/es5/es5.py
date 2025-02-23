def remove_bom(input_file, output_file):
    # Lettura del file di input
    with open(input_file, 'rb') as file:
        content = file.read()

BOM = b'\xef\xbb\xbf'
if content.startswith(BOM):
    print('BOM is retrived')
    content = content[len[BOM]:]
else:
    print("BOM is not retrived")

with open(output_file, 'wb') as file:
    file.write(content)

# Test
input_file = 'Prova.txt'
output_file = 'ProvaSenzaBOM.txt'
print(f"The file is now saved without BOM in {output_file}")