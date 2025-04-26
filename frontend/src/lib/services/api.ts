import type {
	Item,
	User,
	Category,
	RentalRequest,
	SearchParams,
	ItemWithOwner,
	RentalWithDetails
} from '../types';

const API_URL = 'http://localhost:8080';

// Helper function to handle API responses
async function handleResponse<T>(response: Response): Promise<T> {
	if (!response.ok) {
		const errorData = await response.json().catch(() => null);
		throw new Error(errorData?.message || `API error: ${response.status}`);
	}
	return response.json() as Promise<T>;
}

// Common fetch options to use with all API calls
const getCommonOptions = (token: string | null = null): RequestInit => {
	const options: RequestInit = {
		credentials: 'include', // Include credentials like cookies if necessary
		headers: {
			'Content-Type': 'application/json'
		}
	};

	// Add authorization token if available
	if (token) {
		options.headers = {
			...options.headers,
			Authorization: `Bearer ${token}`
		};
	}

	return options;
};

// Helper function to get token
const getToken = (): string | null => {
	return localStorage.getItem('token');
};

// Authentication
export async function login(email: string, password: string) {
	const options = getCommonOptions();
	options.method = 'POST';
	options.body = JSON.stringify({ email, password });

	try {
		const response = await fetch(`${API_URL}/login`, options);
		return handleResponse(response);
	} catch (error) {
		console.error('Login error:', error);
		throw error;
	}
}

// User endpoints
export async function getAllUsers(): Promise<User[]> {
	const token = getToken();
	const options = getCommonOptions(token);

	try {
		const response = await fetch(`${API_URL}/api/users`, options);
		return handleResponse<User[]>(response);
	} catch (error) {
		console.error('Error fetching users:', error);
		throw error;
	}
}

export async function getUser(id: number): Promise<User> {
	const token = getToken();
	const options = getCommonOptions(token);

	try {
		const response = await fetch(`${API_URL}/api/user/${id}`, options);
		return handleResponse<User>(response);
	} catch (error) {
		console.error(`Error fetching user ${id}:`, error);
		throw error;
	}
}

// Item endpoints
export async function getAllItems(): Promise<Item[]> {
	const token = getToken();
	const options = getCommonOptions(token);

	try {
		const response = await fetch(`${API_URL}/api/items`, options);
		return handleResponse<Item[]>(response);
	} catch (error) {
		console.error('Error fetching items:', error);
		throw error;
	}
}

export async function getAvailableItems(): Promise<ItemWithOwner[]> {
	const token = getToken();
	const options = getCommonOptions(token);

	try {
		const response = await fetch(`${API_URL}/api/items/available`, options);
		return handleResponse<ItemWithOwner[]>(response);
	} catch (error) {
		console.error('Error fetching available items:', error);
		throw error;
	}
}

export async function getItem(id: number): Promise<Item> {
	const token = getToken();
	const options = getCommonOptions(token);

	try {
		const response = await fetch(`${API_URL}/api/items/${id}`, options);
		return handleResponse<Item>(response);
	} catch (error) {
		console.error(`Error fetching item ${id}:`, error);
		throw error;
	}
}

export async function createItem(item: Item): Promise<Item> {
	const token = getToken();
	const options = getCommonOptions(token);
	options.method = 'POST';
	options.body = JSON.stringify(item);

	try {
		const response = await fetch(`${API_URL}/api/items`, options);
		return handleResponse<Item>(response);
	} catch (error) {
		console.error('Error creating item:', error);
		throw error;
	}
}

export async function updateItem(id: number, item: Item): Promise<Item> {
	const token = getToken();
	const options = getCommonOptions(token);
	options.method = 'PUT';
	options.body = JSON.stringify(item);

	try {
		const response = await fetch(`${API_URL}/api/items/${id}`, options);
		return handleResponse<Item>(response);
	} catch (error) {
		console.error(`Error updating item ${id}:`, error);
		throw error;
	}
}

export async function deleteItem(id: number): Promise<{ message: string }> {
	const token = getToken();
	const options = getCommonOptions(token);
	options.method = 'DELETE';

	try {
		const response = await fetch(`${API_URL}/api/items/${id}`, options);
		return handleResponse<{ message: string }>(response);
	} catch (error) {
		console.error(`Error deleting item ${id}:`, error);
		throw error;
	}
}

export async function searchItems(params: SearchParams): Promise<Item[]> {
	const searchParams = new URLSearchParams();
	if (params.query) searchParams.append('query', params.query);
	if (params.categoryID !== undefined)
		searchParams.append('category_id', params.categoryID.toString());
	if (params.minPrice !== undefined) searchParams.append('min_price', params.minPrice.toString());
	if (params.maxPrice !== undefined) searchParams.append('max_price', params.maxPrice.toString());
	if (params.available !== undefined) searchParams.append('available', params.available.toString());

	const token = getToken();
	const options = getCommonOptions(token);

	try {
		const response = await fetch(`${API_URL}/api/items/search?${searchParams}`, options);
		return handleResponse<Item[]>(response);
	} catch (error) {
		console.error('Error searching items:', error);
		throw error;
	}
}

// Category endpoints
export async function getAllCategories(): Promise<Category[]> {
	const token = getToken();
	const options = getCommonOptions(token);

	try {
		const response = await fetch(`${API_URL}/api/categories`, options);
		return handleResponse<Category[]>(response);
	} catch (error) {
		console.error('Error fetching categories:', error);
		throw error;
	}
}

// Rental endpoints
export async function createRentalRequest(rental: RentalRequest): Promise<RentalRequest> {
	const token = getToken();
	const options = getCommonOptions(token);
	options.method = 'POST';
	options.body = JSON.stringify(rental);

	try {
		const response = await fetch(`${API_URL}/api/rentals`, options);
		return handleResponse<RentalRequest>(response);
	} catch (error) {
		console.error('Error creating rental request:', error);
		throw error;
	}
}

export async function getMyRentals(): Promise<RentalWithDetails[]> {
	const token = getToken();
	const options = getCommonOptions(token);

	try {
		const response = await fetch(`${API_URL}/api/rentals/my`, options);
		return handleResponse<RentalWithDetails[]>(response);
	} catch (error) {
		console.error('Error fetching my rental requests:', error);
		throw error;
	}
}
