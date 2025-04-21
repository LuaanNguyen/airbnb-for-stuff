<script>
  import { goto } from '$app/navigation';
  
  let email = '';
  let password = '';
  let error = '';
  let loading = false;
  
  // Backend API URL - update this to match your Go server
  const API_URL = 'http://localhost:8080'; // or whatever port your Go server is running on

  async function handleLogin() {
    loading = true;
    error = '';
    
    try {
      const response = await fetch(`${API_URL}/login`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ email, password })
      });
      
      if (!response.ok) {
        const errorData = await response.text();
        throw new Error(errorData || 'Login failed');
      }
      
      const data = await response.json();
      
      // Store JWT token and user info in localStorage
      localStorage.setItem('token', data.token);
      localStorage.setItem('userId', data.userID);
      localStorage.setItem('firstName', data.firstName);
      localStorage.setItem('lastName', data.lastName);
      
      // Redirect to home page
      goto('/items');
    } catch (err) {
      error = err instanceof Error ? err.message : 'An unknown error occurred';
    } finally {
      loading = false;
    }
  }
</script>

<div class="max-w-md mx-auto p-8 rounded-lg shadow-md">
  <h1 class="text-center text-2xl font-bold mb-6">Login</h1>
  
  {#if error}
    <div class="bg-red-50 text-red-800 p-3 rounded-md mb-4">
      {error}
    </div>
  {/if}
  
  <form on:submit|preventDefault={handleLogin}>
    <div class="mb-4">
      <label for="email" class="block mb-2 font-medium">Email</label>
      <input 
        type="email" 
        id="email" 
        bind:value={email} 
        required
        placeholder="Enter your email"
        class="w-full px-3 py-2 border border-gray-300 text-base focus:outline-none focus:ring-2 focus:ring-indigo-500"
      />
    </div>
    
    <div class="mb-4">
      <label for="password" class="block mb-2 font-medium">Password</label>
      <input 
        type="password" 
        id="password" 
        bind:value={password} 
        required
        placeholder="Enter your password"
        class="w-full px-3 py-2 border border-gray-300 text-base focus:outline-none focus:ring-2 focus:ring-indigo-500"
      />
    </div>
    
    <button 
      type="submit" 
      disabled={loading}
      class="w-full py-3 px-4 hover:bg-black hover:text-white border disabled:cursor-not-allowed"
    >
      {loading ? 'Logging in...' : 'Login'}
    </button>
  </form>
</div>
