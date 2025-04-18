## Run Svelte

```
cd frontend
bun run dev --open
```

Check `localhost:5173/`

## Run server

Create a `.env` file in `/backend` with:

```
POSTGRES_URL=postgres://<username>:<password>@<host>:<port>/<dbname>?sslmode=require
```

Run the program

```
cd backend
go mod tidy # install depencies
go run main.go
```

Check `localhost:8080/`

## Generate mock data

```
cd data
python -m venv venv
source venv/bin/activate
pip install -r requirements.txt
python script.py
```

Mock data should be generated in `/fake_data_csv`
