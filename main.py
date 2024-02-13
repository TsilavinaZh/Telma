import json

def generate_numbers(debut, fin):
    numbers = []
    for i in range(debut, fin + 1):
        numbers.append("0" + str(i).zfill(10))
    return numbers

telma34 = generate_numbers(3400000000, 3499999999)
telma38 = generate_numbers(3800000000, 3899999999)

data = {
    "telma34": telma34,
    "telma38": telma38
}

json_data = json.dumps(data, indent=2)

with open("Telma.json", "w") as file:
    file.write(json_data)

print("success.")