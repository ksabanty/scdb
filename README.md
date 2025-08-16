# SCDB - Sports Card Database

A simple, extensible CLI application to manage your sports card collection.  
Built with Go and `cobra`.

---

## 🚀 Features
- Add a card to your collection
- Delete a card from your collection (by player name)
- List all cards in a formatted table
- Data stored locally in JSON (`data/collection.json`)
- Extensible project structure for future features

---

## 📦 Requirements
- Go (latest version recommended, e.g. 1.22+)

---

## 🔧 Installation

Clone this repository:

```bash
git clone https://github.com/ksabanty/scdb.git
cd scdb
```

---

## ⚙️ Build & Install Libraries

Before building, install required libraries:

```bash
go get github.com/olekukonko/tablewriter
go get github.com/spf13/cobra
```

To build the project:

```bash
go build -o scdb
```

---

## 📝 Usage

After building the project, you can use the following commands:

### Add a card
```bash
scdb add
```
Interactively add a new card to your collection.

### List all cards
```bash
scdb list
```
Display all cards in your collection in a table format.

### Delete a card
```bash
scdb delete --player "Lebron James" --team "Lakers" --year "2020"
```
You can specify additional fields (such as `--team`, `--manufacturer`, `--year`, etc.) to uniquely identify and delete the correct card if a player has multiple cards in your collection. At minimum, `--player` is required.

Example:
```bash
scdb delete --player "Lebron James" --team "Lakers" --year "2020" --rookie
```
This will delete the card matching all provided criteria.

---

## 📄 License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---
