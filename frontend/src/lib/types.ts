export interface Item {
	id?: number;
	name: string;
	description: string;
	category_id: number;
	owner_id?: number;
	price: number;
	quantity: number;
	available: boolean;
	image?: string;
	date_listed?: Date;
}

export interface User {
	id: number;
	email: string;
	phone_number: string;
	first_name: string;
	last_name: string;
	nick_name?: string;
}

export interface Category {
	id: number;
	name: string;
	description: string;
}

export interface RentalRequest {
	id?: number;
	item_id: number;
	renter_id?: number;
	start_date: string;
	end_date: string;
	status?: string;
	total_price: number;
}

export interface SearchParams {
	query?: string;
	categoryID?: number;
	minPrice?: number;
	maxPrice?: number;
	available?: boolean;
}

export interface ItemWithOwner {
	id: number;
	name: string;
	description: string;
	price: number;
	owner_id: number;
	owner_name: string;
	available: boolean;
	image?: string;
}

export interface LoginResponse {
	token: string;
	user_id: number;
	first_name: string;
	last_name: string;
}

export interface RentalWithDetails {
	id: number;
	item_id: number;
	item_name: string;
	description: string;
	start_date: string;
	end_date: string;
	status: string;
	total_price: number;
	owner_name: string;
}
