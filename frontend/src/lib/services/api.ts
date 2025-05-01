import type {
	Item,
	User,
	Category,
	RentalRequest,
	SearchParams,
	ItemWithOwner,
	RentalWithDetails,
	LoginResponse
} from '../types';

// --- MOCK DATA ---

// Mock Users (replace with more realistic data if needed)
const mockUsers: User[] = [
	{
		id: 1,
		email: 'alice@example.com',
		phone_number: '123-456-7890',
		first_name: 'Alice',
		last_name: 'Smith'
	},
	{
		id: 2,
		email: 'bob@example.com',
		phone_number: '987-654-3210',
		first_name: 'Bob',
		last_name: 'Jones'
	}
];

// Mock Categories
const mockCategories: Category[] = [
	{ id: 1, name: 'Gardening Tools', description: 'Tools for your garden' },
	{ id: 2, name: 'Camping Gear', description: 'Equipment for outdoor adventures' },
	{ id: 3, name: 'Power Tools', description: 'Tools for construction and DIY' },
	{ id: 4, name: 'Kitchen Appliances', description: 'Small appliances for your kitchen' }
];
// let nextCategoryId = 5; // Commented out: Unused variable

// Mock Items
// Note: Price is in cents
const mockItems: Item[] = [
	{
		id: 1,
		name: 'Electric Lawn Mower',
		description: 'Powerful electric lawn mower, great for medium yards.',
		category_id: 1,
		owner_id: 1,
		price: 5000,
		quantity: 1,
		available: true,
		date_listed: new Date('2023-10-01'),
		image: 'https://placehold.co/600x400/png?text=Lawn+Mower'
	},
	{
		id: 2,
		name: '4-Person Camping Tent',
		description: 'Spacious and waterproof tent.',
		category_id: 2,
		owner_id: 2,
		price: 2500,
		quantity: 1,
		available: true,
		date_listed: new Date('2023-10-05'),
		image: 'https://placehold.co/600x400/png?text=Camping+Tent'
	},
	{
		id: 3,
		name: 'Cordless Power Drill',
		description: '18V cordless drill with two batteries and charger.',
		category_id: 3,
		owner_id: 1,
		price: 1500,
		quantity: 1,
		available: false,
		date_listed: new Date('2023-10-10'),
		image: 'https://placehold.co/600x400/png?text=Power+Drill'
	},
	{
		id: 4,
		name: 'High-Speed Blender',
		description: 'Blender for smoothies and more.',
		category_id: 4,
		owner_id: 2,
		price: 3000,
		quantity: 1,
		available: true,
		date_listed: new Date('2023-10-15'),
		image: 'https://placehold.co/600x400/png?text=Blender'
	},
	{
		id: 5,
		name: 'Pressure Washer',
		description: 'Electric pressure washer for cleaning decks, patios, and cars.',
		category_id: 3,
		owner_id: 1,
		price: 4500,
		quantity: 1,
		available: true,
		date_listed: new Date('2023-11-01'),
		image: 'https://placehold.co/600x400/png?text=Pressure+Washer'
	},
	{
		id: 6,
		name: 'Stand Mixer',
		description: 'Heavy-duty stand mixer for baking enthusiasts.',
		category_id: 4,
		owner_id: 2,
		price: 3500,
		quantity: 1,
		available: true,
		date_listed: new Date('2023-11-05'),
		image: 'https://placehold.co/600x400/png?text=Stand+Mixer'
	},
	{
		id: 7,
		name: 'Hedge Trimmer',
		description: 'Electric hedge trimmer for shaping shrubs and bushes.',
		category_id: 1,
		owner_id: 1,
		price: 2000,
		quantity: 2,
		available: true,
		date_listed: new Date('2023-11-10'),
		image: 'https://placehold.co/600x400/png?text=Hedge+Trimmer'
	},
	{
		id: 8,
		name: 'Backpack (Hiking)',
		description: '60L hiking backpack, suitable for multi-day trips.',
		category_id: 2,
		owner_id: 2,
		price: 1800,
		quantity: 1,
		available: false,
		date_listed: new Date('2023-11-15'),
		image: 'https://placehold.co/600x400/png?text=Hiking+Backpack'
	},
	{
		id: 9,
		name: 'Air Fryer',
		description: '5-quart air fryer for healthier cooking.',
		category_id: 4,
		owner_id: 1,
		price: 2800,
		quantity: 1,
		available: true,
		date_listed: new Date('2023-11-20'),
		image: 'https://placehold.co/600x400/png?text=Air+Fryer'
	},
	{
		id: 10,
		name: 'Sleeping Bag (-5C)',
		description: 'Warm sleeping bag rated for -5 Celsius.',
		category_id: 2,
		owner_id: 2,
		price: 1200,
		quantity: 3,
		available: true,
		date_listed: new Date('2023-11-22'),
		image: 'https://placehold.co/600x400/png?text=Sleeping+Bag'
	},
	{
		id: 11,
		name: 'Circular Saw',
		description: '7-1/4 inch circular saw with laser guide.',
		category_id: 3,
		owner_id: 1,
		price: 3200,
		quantity: 1,
		available: true,
		date_listed: new Date('2023-11-25'),
		image: 'https://placehold.co/600x400/png?text=Circular+Saw'
	},
	{
		id: 12,
		name: 'Leaf Blower (Electric)',
		description: 'Corded electric leaf blower, lightweight.',
		category_id: 1,
		owner_id: 2,
		price: 1600,
		quantity: 1,
		available: true,
		date_listed: new Date('2023-11-28'),
		image: 'https://placehold.co/600x400/png?text=Leaf+Blower'
	}
];
let nextItemId = 13;

// Mock Rentals
// Commented out: Unused variable
// const mockRentals: RentalWithDetails[] = [
// 	// Add mock rental data if needed for getMyRentals
// ];
let nextRentalId = 1;

// --- MOCK HELPER FUNCTIONS ---

// Simulate network delay
const delay = (ms: number) => new Promise((resolve) => setTimeout(resolve, ms));

// Get user name (replace with actual user lookup if needed)
const getUserName = (userId: number | undefined): string => {
	if (userId === undefined) return 'Unknown User';
	const user = mockUsers.find((u) => u.id === userId);
	return user ? `${user.first_name} ${user.last_name}` : 'Unknown User';
};

// --- MOCKED API FUNCTIONS ---

// Authentication
export async function login(email: string, _password: string): Promise<LoginResponse> {
	console.log('SIMULATED: Login attempt for', email);
	await delay(500);
	const user = mockUsers.find((u) => u.email === email); // Simple mock: find user by email, ignore password
	if (user) {
		const token = `mock-token-for-${user.id}-${Date.now()}`;
		localStorage.setItem('token', token); // Store mock token
		localStorage.setItem('user_id', user.id.toString()); // Store user ID
		localStorage.setItem('user_name', `${user.first_name} ${user.last_name}`);
		console.log('SIMULATED: Login successful for', user.first_name);
		return {
			token: token,
			user_id: user.id,
			first_name: user.first_name,
			last_name: user.last_name
		};
	} else {
		console.log('SIMULATED: Login failed for', email);
		throw new Error('Invalid credentials (mock)');
	}
}

// Mock token retrieval - bypasses actual token checking for other functions
// Commented out: Unused function
// const getToken = (): string | null => {
// 	// In a real app, you'd validate this token
// 	return localStorage.getItem('token');
// };

// User endpoints (basic mocks)
export async function getAllUsers(): Promise<User[]> {
	console.log('SIMULATED: Fetching all users');
	await delay(200);
	return [...mockUsers]; // Return a copy
}

export async function getUser(id: number): Promise<User> {
	console.log(`SIMULATED: Fetching user ${id}`);
	await delay(150);
	const user = mockUsers.find((u) => u.id === id);
	if (!user) {
		throw new Error(`User with id ${id} not found (mock)`);
	}
	return { ...user }; // Return a copy
}

// Item endpoints
export async function getAllItems(): Promise<Item[]> {
	console.log('SIMULATED: Fetching all items');
	await delay(300);
	return [...mockItems]; // Return a copy
}

export async function getAvailableItems(): Promise<ItemWithOwner[]> {
	console.log('SIMULATED: Fetching available items');
	await delay(300);
	const available = mockItems.filter((item) => item.available);
	// Map to ItemWithOwner, adding owner name
	return available.map((item) => ({
		id: item.id!,
		name: item.name,
		description: item.description,
		price: item.price,
		owner_id: item.owner_id!,
		owner_name: getUserName(item.owner_id),
		available: item.available,
		image: item.image
	}));
}

export async function getItem(id: number): Promise<Item> {
	console.log(`SIMULATED: Fetching item ${id}`);
	await delay(200);
	const item = mockItems.find((item) => item.id === id);
	if (!item) {
		throw new Error(`Item with id ${id} not found (mock)`);
	}
	return { ...item }; // Return a copy
}

// Helper to get current logged-in user ID (mocked)
const getCurrentUserId = (): number | undefined => {
	const userIdStr = localStorage.getItem('user_id');
	return userIdStr ? parseInt(userIdStr, 10) : undefined;
};

export async function createItem(
	itemData: Omit<Item, 'id' | 'owner_id' | 'date_listed' | 'available'>
): Promise<Item> {
	console.log('SIMULATED: Creating item', itemData);
	await delay(400);
	const currentUserId = getCurrentUserId();
	if (!currentUserId) {
		throw new Error('User not logged in (mock)'); // Simulate auth check
	}

	const newItem: Item = {
		...itemData,
		id: nextItemId++,
		owner_id: currentUserId, // Assign to current mock user
		available: true, // Default to available
		quantity: itemData.quantity || 1, // Default quantity
		date_listed: new Date()
	};
	mockItems.push(newItem);
	console.log('SIMULATED: Item created', newItem);
	return { ...newItem }; // Return a copy
}

export async function updateItem(
	id: number,
	itemData: Partial<Omit<Item, 'id' | 'owner_id' | 'date_listed'>>
): Promise<Item> {
	console.log(`SIMULATED: Updating item ${id}`, itemData);
	await delay(350);
	const itemIndex = mockItems.findIndex((item) => item.id === id);
	if (itemIndex === -1) {
		throw new Error(`Item with id ${id} not found for update (mock)`);
	}
	// Ensure owner cannot be changed via this method
	const currentUserId = getCurrentUserId();
	if (mockItems[itemIndex].owner_id !== currentUserId) {
		throw new Error('Cannot update item owned by another user (mock)');
	}

	mockItems[itemIndex] = {
		...mockItems[itemIndex],
		...itemData,
		// Ensure read-only fields are not overwritten from itemData
		id: mockItems[itemIndex].id,
		owner_id: mockItems[itemIndex].owner_id,
		date_listed: mockItems[itemIndex].date_listed
	};
	console.log('SIMULATED: Item updated', mockItems[itemIndex]);
	return { ...mockItems[itemIndex] }; // Return a copy
}

export async function deleteItem(id: number): Promise<{ message: string }> {
	console.log(`SIMULATED: Deleting item ${id}`);
	await delay(500);
	const itemIndex = mockItems.findIndex((item) => item.id === id);
	if (itemIndex === -1) {
		throw new Error(`Item with id ${id} not found for deletion (mock)`);
	}
	// Ensure only owner can delete
	const currentUserId = getCurrentUserId();
	if (mockItems[itemIndex].owner_id !== currentUserId) {
		throw new Error('Cannot delete item owned by another user (mock)');
	}

	mockItems.splice(itemIndex, 1);
	console.log('SIMULATED: Item deleted successfully');
	return { message: 'Item deleted successfully (mock)' };
}

// Extended Search Logic
export async function searchItems(params: SearchParams): Promise<ItemWithOwner[]> {
	console.log('SIMULATED: Searching items with params:', params);
	await delay(300);

	let results = [...mockItems]; // Start with a copy

	if (params.available !== undefined) {
		results = results.filter((item) => item.available === params.available);
	}
	if (params.query) {
		const queryLower = params.query.toLowerCase();
		results = results.filter(
			(item) =>
				item.name.toLowerCase().includes(queryLower) ||
				item.description.toLowerCase().includes(queryLower)
		);
	}
	if (params.categoryID !== undefined && params.categoryID !== null) {
		results = results.filter((item) => item.category_id === params.categoryID);
	}
	if (params.minPrice !== undefined && params.minPrice !== null) {
		// Convert filter price ($) to cents for comparison
		results = results.filter((item) => item.price >= params.minPrice! * 100);
	}
	if (params.maxPrice !== undefined && params.maxPrice !== null) {
		// Convert filter price ($) to cents for comparison
		results = results.filter((item) => item.price <= params.maxPrice! * 100);
	}

	// Map results to ItemWithOwner
	return results.map((item) => ({
		id: item.id!,
		name: item.name,
		description: item.description,
		price: item.price,
		owner_id: item.owner_id!,
		owner_name: getUserName(item.owner_id),
		available: item.available,
		image: item.image
	}));
}

// Category endpoints
export async function getAllCategories(): Promise<Category[]> {
	console.log('SIMULATED: Fetching all categories');
	await delay(200);
	return [...mockCategories]; // Return a copy
}

// Rental endpoints (basic mocks)
export async function createRentalRequest(
	rentalData: Omit<RentalRequest, 'id' | 'renter_id' | 'status'>
): Promise<RentalRequest> {
	console.log('SIMULATED: Creating rental request', rentalData);
	await delay(450);
	const currentUserId = getCurrentUserId();
	if (!currentUserId) {
		throw new Error('User not logged in (mock)');
	}
	const item = mockItems.find((i) => i.id === rentalData.item_id);
	if (!item) {
		throw new Error(`Item ${rentalData.item_id} not found (mock)`);
	}
	if (!item.available) {
		throw new Error(`Item ${rentalData.item_id} is not available for rent (mock)`);
	}

	const newRental: RentalRequest = {
		...rentalData,
		id: nextRentalId++,
		renter_id: currentUserId,
		status: 'pending' // Default status
	};
	// Note: In a real app, you'd store this rental request. We are not storing it in this basic mock.
	console.log('SIMULATED: Rental request created', newRental);

	// Mark item as unavailable (simple simulation) - A real app needs date checking
	const itemIndex = mockItems.findIndex((i) => i.id === rentalData.item_id);
	if (itemIndex > -1) {
		mockItems[itemIndex].available = false;
	}

	return { ...newRental }; // Return a copy
}

export async function getMyRentals(): Promise<RentalWithDetails[]> {
	console.log('SIMULATED: Fetching my rentals');
	await delay(300);
	const currentUserId = getCurrentUserId();
	if (!currentUserId) {
		return []; // Return empty if not logged in
	}
	// This is a placeholder - requires storing created rentals or adding more mock data
	// Returning a predefined list for now if needed, otherwise empty
	// Example: return mockRentals.filter(r => r.renter_id === currentUserId);
	console.log('SIMULATED: Returning mock rentals (currently empty array)');
	return []; // Return empty array as we aren't storing created rentals yet
}

// Get items owned by the current user
export async function getMyItems(): Promise<Item[]> {
	console.log('SIMULATED: Fetching my items');
	await delay(250);
	const currentUserId = getCurrentUserId();
	if (!currentUserId) {
		console.warn('SIMULATED: getMyItems called without logged-in user.');
		return []; // Return empty if no user is logged in
	}
	const userItems = mockItems.filter((item) => item.owner_id === currentUserId);
	return [...userItems]; // Return a copy
}
