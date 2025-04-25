// src/lib/auth.ts
import { writable, type Writable } from 'svelte/store';
import type { User, LoginResponse } from './types';
import { browser } from '$app/environment';

// Create stores for authentication
export const user: Writable<User | null> = writable(null);
export const token: Writable<string | null> = writable(null);
export const isAuthenticated: Writable<boolean> = writable(false);

// Initialize from localStorage (only runs in browser)
if (browser) {
	const storedToken = localStorage.getItem('token');
	const storedUser = localStorage.getItem('user');

	if (storedToken) {
		token.set(storedToken);
		isAuthenticated.set(true);
	}

	if (storedUser) {
		try {
			const parsedUser = JSON.parse(storedUser);
			user.set(parsedUser);
		} catch (e) {
			console.error('Failed to parse stored user data', e);
		}
	}
}

// Helper functions
export function setAuth(loginResponse: LoginResponse): void {
	// Set the token and user in localStorage
	if (browser) {
		localStorage.setItem('token', loginResponse.token);
		const userData: User = {
			id: loginResponse.user_id,
			first_name: loginResponse.first_name,
			last_name: loginResponse.last_name,
			email: '', // Backend doesn't return this in login response
			phone_number: '' // Backend doesn't return this in login response
		};
		localStorage.setItem('user', JSON.stringify(userData));

		// Update stores
		token.set(loginResponse.token);
		user.set(userData);
		isAuthenticated.set(true);
	}
}

export function clearAuth(): void {
	// Clear from localStorage
	if (browser) {
		localStorage.removeItem('token');
		localStorage.removeItem('user');
	}

	// Reset stores
	token.set(null);
	user.set(null);
	isAuthenticated.set(false);
}

export function getToken(): string | null {
	if (browser) {
		return localStorage.getItem('token');
	}
	return null;
}
