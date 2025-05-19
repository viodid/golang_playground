def main():
    out = []
    for i in range(100_000_000):
        out.append({
            "name": "Bob",
            "surname": "Smith",
            "age": 42,
            })
    print(out[:1])

if __name__ == "__main__":
    main()

