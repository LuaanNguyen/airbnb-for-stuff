### Run Svelte

```
cd frontend
bun run dev --open
```

Check `localhost:5173/`

### Run the server

```
cd backend
go mod tidy # install depencies
go run main.go
```

Check `localhost:8080/`

### Generate mock data

```
cd mock-data
python -m venv venv
source venv/bin/activate
pip install -r requirements.txt
python script.py
```

Mock data should be generated in `/fake_data_csv`
