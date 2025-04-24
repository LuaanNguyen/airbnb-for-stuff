# AirBnb For Stuff

Built with [Svelte](https://svelte.dev/), [GO](https://go.dev/), [PostgresSQL](https://www.postgresql.org/).

This project is an simple marketplace that allows users to rent out various personal items.

## App Structure

```
├── backend
│   ├── main.go
|   ├── go.mod
|   ├── go.sum
|   ├── .env
│   ├── db                  // for DB connections
│   │   ├── db.go.go
│   │   ├── queries.sql
│   │   ├── schema.sql      // DB schema
│   ├── handlers          // API core handlers
│   │   ├── handlers.go
|   ├── middleware          // auth, CORS
│   │   ├── auth.go
│   │   ├── cors.go
│   └── models
│   |    └── model_functions.go     // DB functions
│   |    └── ...                    // Models for our application
|   |
├── frontend    //svelte app
├── data    // mock data generation
└── backup  // application's data
```

## Run Svelte app

```
cd frontend
npm i
npm run dev
npm run build
npm run preview
```

Check `localhost:5173/`

## Run GO app

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

### 💿 Backup data are stored in `data/`
