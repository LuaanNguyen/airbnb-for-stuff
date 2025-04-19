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

## TODOS

User Management

- [ ] `/api/user/register` - For new user registration
- [ ] `/api/user/login` - For user authentication
- [ ] `/api/user/{id}` - PUT endpoint for updating user profile

Item Management

- [ ] Full CRUD operations
- [ ] `/api/items/search` - For searching items with filters
- [ ] `/api/items/available` - For getting available items
- [ ] `/api/items/{id}` - For individual items operations

Category Management:

- [ ] `/api/categories` - Get all categories
- [ ] `/api/categories/{id}/items` - Get items by category

Transaction Management

- [ ] `/api/transactions` - Create new transactions
- [ ] `/api/user/{id}/transactions` - Get user's transactiopn history
- [ ] `/api/transactions/{id}` - Get and update specific transactions

Review Systems

- [ ] `/api/reviews` - Create new reviews
- [ ] `/api/items/{id}/reviews` - Get reviews for an item
- [ ] `/api/user/{id}/reviews` - Get reviews by a user
