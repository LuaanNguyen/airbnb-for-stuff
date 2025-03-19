from faker import Faker
import random
from datetime import datetime, timedelta
import csv
from pathlib import Path

# Initialize Faker
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
        # Convert boolean values to lowercase strings for PostgreSQL
        for row in data:
            for key, value in row.items():
                if isinstance(value, bool):
                    row[key] = str(value).lower()
            writer.writerow(row)
    print(f"Created {filename} with {len(data)} records")

def generate_users(num_users=500):
    users = []
    used_emails = set()  # Keep track of used emails
    
    for i in range(num_users):
        # Keep generating emails until we get a unique one
        while True:
            email = fake.unique.email()  # Use Faker's unique provider
            if email not in used_emails:
                used_emails.add(email)
                break
        
        # Generate a phone number in format: ###-###-####
        phone = fake.numerify('###-###-####')
        
        user = {
            'u_id': i + 1,
            'u_email': email,
            'u_phone_number': phone,
            'u_first_name': fake.first_name(),
            'u_last_name': fake.last_name(),
            'u_nick_name': fake.user_name() if random.choice([True, False]) else None,
            'u_password': fake.password(length=12)
        }
        users.append(user)
    
    save_to_csv(users, 'users.csv', 
                ['u_id', 'u_email', 'u_phone_number', 'u_first_name', 
                 'u_last_name', 'u_nick_name', 'u_password'])
    return users

def generate_addresses(users, addresses_per_user=2):
    addresses = []
    address_id = 1
    
    for user in users:
        for _ in range(random.randint(1, addresses_per_user)):
            address = {
                'a_id': address_id,
                'u_id': user['u_id'],
                'a_street': fake.street_address()[:255],  # Limit street address length
                'a_city': fake.city()[:100],  # Limit city length
                'a_state': fake.state()[:100],  # Limit state length
                'a_zipcode': fake.zipcode()[:20],  # Limit zipcode length
                'a_country': fake.country()[:100]  # Limit country length
            }
            addresses.append(address)
            address_id += 1
    
    save_to_csv(addresses, 'addresses.csv',
                ['a_id', 'u_id', 'a_street', 'a_city', 'a_state', 
                 'a_zipcode', 'a_country'])
    return addresses

def generate_categories():
    categories = [
        {'c_id': 1, 'c_name': 'Electronics', 'c_description': 'Electronic devices and accessories'},
        {'c_id': 2, 'c_name': 'Outdoor Equipment', 'c_description': 'Camping and hiking gear'},
        {'c_id': 3, 'c_name': 'Tools', 'c_description': 'Power and hand tools'},
        {'c_id': 4, 'c_name': 'Sports Equipment', 'c_description': 'Sports and fitness gear'},
        {'c_id': 5, 'c_name': 'Musical Instruments', 'c_description': 'Instruments and audio equipment'},
        {'c_id': 6, 'c_name': 'Photography', 'c_description': 'Cameras and accessories'},
        {'c_id': 7, 'c_name': 'Party Supplies', 'c_description': 'Party decorations and equipment'},
        {'c_id': 8, 'c_name': 'Books', 'c_description': 'Books and reading materials'},
        {'c_id': 9, 'c_name': 'Gaming', 'c_description': 'Video games and consoles'},
        {'c_id': 10, 'c_name': 'Home & Garden', 'c_description': 'Home improvement and gardening tools'},
        {'c_id': 11, 'c_name': 'Vehicles', 'c_description': 'Cars, bikes, and other vehicles'},
        {'c_id': 12, 'c_name': 'Fashion', 'c_description': 'Clothing and accessories'}
    ]
    
    save_to_csv(categories, 'categories.csv', ['c_id', 'c_name', 'c_description'])
    return categories

def generate_items(users, categories, num_items=1000):
    items = []
    for i in range(num_items):
        category = random.choice(categories)
        date_listed = fake.date_time_between(start_date='-1y').isoformat()
        
        item = {
            'i_id': i + 1,
            'i_name': fake.catch_phrase()[:255],  # Limit name length
            'i_description': fake.text(max_nb_chars=200),  # Already limited
            'i_image': None,  # Binary data would need special handling
            'c_id': category['c_id'],
            'i_price': random.randint(500, 20000),  # $5 to $200 (in cents)
            'i_date_listed': date_listed,
            'i_quantity': random.randint(1, 10),
            'i_available': random.choice([True, True, False])  # 66% available
        }
        items.append(item)
    
    save_to_csv(items, 'items.csv',
                ['i_id', 'i_name', 'i_description', 'i_image', 'c_id',
                 'i_price', 'i_date_listed', 'i_quantity', 'i_available'])
    return items

def generate_transactions(users, items, num_transactions=2000):
    transaction_types = ['Purchase', 'Sale', 'Refund', 'Rental']
    transactions = []
    for i in range(num_transactions):
        item = random.choice(items)
        user = random.choice(users)
        transaction_date = fake.date_time_between(
            start_date=datetime.fromisoformat(item['i_date_listed']),
            end_date='now'
        ).isoformat()
        
        transaction = {
            't_id': i + 1,
            'u_id': user['u_id'],
            't_type': random.choice(transaction_types),
            'i_id': item['i_id'],
            't_date': transaction_date,
            't_amount': item['i_price'] * random.randint(1, item['i_quantity'])
        }
        transactions.append(transaction)
    
    save_to_csv(transactions, 'transactions.csv',
                ['t_id', 'u_id', 't_type', 'i_id', 't_date', 't_amount'])
    return transactions

def generate_reviews(users, num_reviews=1500):
    reviews = []
    for i in range(num_reviews):
        user = random.choice(users)
        
        review = {
            'r_id': i + 1,
            'r_comment': fake.text(max_nb_chars=150),  # Limit comment length
            'r_star': random.randint(1, 5),
            'u_id': user['u_id']
        }
        reviews.append(review)
    
    save_to_csv(reviews, 'reviews.csv',
                ['r_id', 'r_comment', 'r_star', 'u_id'])
    return reviews

def main():
    print("Starting data generation...")
    users = generate_users()
    addresses = generate_addresses(users)
    categories = generate_categories()
    items = generate_items(users, categories)
    transactions = generate_transactions(users, items)
    reviews = generate_reviews(users)
    print("\nData generation complete! New quantities:")
    print(f"Users: 500 (10x)")
    print(f"Addresses: ~750 (10x)")
    print(f"Categories: 12 (1.5x)")
    print(f"Items: 1000 (10x)")
    print(f"Transactions: 2000 (10x)")
    print(f"Reviews: 1500 (10x)")

if __name__ == "__main__":
    main()
