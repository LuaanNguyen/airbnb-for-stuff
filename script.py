from faker import Faker
import random
from datetime import datetime, timedelta
import csv
from pathlib import Path

# Initialize Faker and create output directory
fake = Faker()
random.seed(42)
Faker.seed(42)

# Create output directory for CSV files
output_dir = Path('fake_data_csv')
output_dir.mkdir(exist_ok=True)

def save_to_csv(data, filename, fieldnames):
    filepath = output_dir / filename
    with open(filepath, 'w', newline='', encoding='utf-8') as f:
        writer = csv.DictWriter(f, fieldnames=fieldnames)
        writer.writeheader()
        writer.writerows(data)
    print(f"Created {filename}")

def generate_users(num_users=50):
    users = []
    for _ in range(num_users):
        user = {
            'id': _ + 1,
            'email': fake.email(),
            'phone_number': fake.phone_number(),
            'first_name': fake.first_name(),
            'last_name': fake.last_name(),
            'date_joined': fake.date_time_between(start_date='-2y').isoformat(),
            'is_active': random.choice([True, True, True, False])
        }
        users.append(user)
    
    save_to_csv(users, 'users.csv', 
                ['id', 'email', 'phone_number', 'first_name', 'last_name', 
                 'date_joined', 'is_active'])
    return users

def generate_addresses(users, addresses_per_user=1):
    addresses = []
    address_id = 1
    
    for user in users:
        for _ in range(random.randint(1, addresses_per_user)):
            address = {
                'id': address_id,
                'user_id': user['id'],
                'street_address': fake.street_address(),
                'city': fake.city(),
                'state': fake.state(),
                'postal_code': fake.postcode(),
                'country': fake.country(),
                'is_default': True if _ == 0 else False
            }
            addresses.append(address)
            address_id += 1
    
    save_to_csv(addresses, 'addresses.csv',
                ['id', 'user_id', 'street_address', 'city', 'state', 
                 'postal_code', 'country', 'is_default'])
    return addresses

def generate_categories():
    categories = [
        {'id': 1, 'name': 'Electronics'},
        {'id': 2, 'name': 'Outdoor Equipment'},
        {'id': 3, 'name': 'Tools'},
        {'id': 4, 'name': 'Sports Equipment'},
        {'id': 5, 'name': 'Musical Instruments'},
        {'id': 6, 'name': 'Photography'},
        {'id': 7, 'name': 'Party Supplies'},
        {'id': 8, 'name': 'Books'},
    ]
    
    save_to_csv(categories, 'categories.csv', ['id', 'name'])
    return categories

def generate_items(users, num_items=100):
    items = []
    for i in range(num_items):
        owner = random.choice(users)
        item = {
            'id': i + 1,
            'owner_id': owner['id'],
            'title': fake.catch_phrase(),
            'description': fake.text(max_nb_chars=200),
            'daily_rate': round(random.uniform(5.0, 200.0), 2),
            'created_at': fake.date_time_between(start_date='-1y').isoformat(),
            'condition': random.choice(['New', 'Like New', 'Good', 'Fair']),
            'is_available': random.choice([True, True, False])
        }
        items.append(item)
    
    save_to_csv(items, 'items.csv',
                ['id', 'owner_id', 'title', 'description', 'daily_rate',
                 'created_at', 'condition', 'is_available'])
    return items

def generate_item_categories(items, categories):
    item_categories = []
    for item in items:
        num_categories = random.randint(1, 3)
        selected_categories = random.sample(categories, num_categories)
        for category in selected_categories:
            item_category = {
                'item_id': item['id'],
                'category_id': category['id']
            }
            item_categories.append(item_category)
    
    save_to_csv(item_categories, 'item_categories.csv',
                ['item_id', 'category_id'])
    return item_categories

def generate_transactions(users, items, num_transactions=200):
    transactions = []
    for i in range(num_transactions):
        item = random.choice(items)
        renter = random.choice([u for u in users if u['id'] != item['owner_id']])
        start_date = fake.date_time_between(start_date='-6m')
        end_date = start_date + timedelta(days=random.randint(1, 14))
        
        transaction = {
            'id': i + 1,
            'item_id': item['id'],
            'renter_id': renter['id'],
            'start_date': start_date.isoformat(),
            'end_date': end_date.isoformat(),
            'total_amount': round(item['daily_rate'] * (end_date - start_date).days, 2),
            'status': random.choice(['Completed', 'Pending', 'Cancelled', 'Active']),
            'created_at': (start_date - timedelta(days=random.randint(1, 7))).isoformat()
        }
        transactions.append(transaction)
    
    save_to_csv(transactions, 'transactions.csv',
                ['id', 'item_id', 'renter_id', 'start_date', 'end_date',
                 'total_amount', 'status', 'created_at'])
    return transactions

def generate_reviews(transactions, num_reviews=150):
    reviews = []
    for i in range(num_reviews):
        transaction = random.choice(transactions)
        review = {
            'id': i + 1,
            'transaction_id': transaction['id'],
            'item_id': transaction['item_id'],
            'reviewer_id': transaction['renter_id'],
            'rating': random.randint(1, 5),
            'comment': fake.text(max_nb_chars=150),
            'created_at': (datetime.fromisoformat(transaction['end_date']) + 
                         timedelta(days=random.randint(1, 5))).isoformat()
        }
        reviews.append(review)
    
    save_to_csv(reviews, 'reviews.csv',
                ['id', 'transaction_id', 'item_id', 'reviewer_id',
                 'rating', 'comment', 'created_at'])
    return reviews

def main():
    users = generate_users()
    addresses = generate_addresses(users)
    categories = generate_categories()
    items = generate_items(users)
    item_categories = generate_item_categories(items, categories)
    transactions = generate_transactions(users, items)
    reviews = generate_reviews(transactions)

if __name__ == "__main__":
    main()
