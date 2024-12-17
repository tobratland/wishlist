<script>
    import { goto } from '$app/navigation';
  </script>
  
  <div class="min-h-screen flex flex-col items-center justify-center bg-snowfall">
    <h1 class="text-5xl text-christmasRed mb-6">🎄 Welcome to the Christmas Wishlist 🎅</h1>
    <p class="text-xl text-gray-800 mb-8">Share your festive wishes with friends and family!</p>
    <div class="flex space-x-4">
      <button
        class="bg-christmasGreen text-white px-6 py-3 rounded shadow"
        on:click={() => goto('/auth/login')}
      >
        🎁 Login
      </button>
      <button
        class="bg-christmasRed text-white px-6 py-3 rounded shadow"
        on:click={() => goto('/auth/register')}
      >
        ✨ Register
      </button>
    </div>
  </div>
  