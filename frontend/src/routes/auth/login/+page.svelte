<script>
    import { goto } from '$app/navigation';
    import { login } from '$lib/api';
  
    let email = '';
    let password = '';
    let error = '';
  
    async function handleLogin() {
      error = '';
      const result = await login(email, password);
      if (result.success) {
        goto('/wishlists');
      } else {
        error = result.message || 'Login failed. Please try again.';
      }
    }
  </script>
  
  <div class="min-h-screen bg-snowfall flex flex-col items-center justify-center">
    <h1 class="text-3xl text-christmasRed mb-4">🎄 Login 🎁</h1>
  
    {#if error}
      <p class="text-red-500">{error}</p>
    {/if}
  
    <form class="space-y-4" on:submit|preventDefault={handleLogin}>
      <input
        placeholder="Email"
        bind:value={email}
        class="block p-2 border rounded"
        required
      />
      <input
        type="password"
        placeholder="Password"
        bind:value={password}
        class="block p-2 border rounded"
        required
      />
      <button class="bg-christmasGreen text-white px-4 py-2">Login</button>
    </form>
  </div>
  