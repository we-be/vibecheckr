import PocketBase from 'https://unpkg.com/pocketbase?module';

// Use HTTPS in the PocketBase URL
const pb = new PocketBase('https://we-be.xyz');

document.getElementById('login-button').addEventListener('click', async () => {
    try {
        const authData = await pb.collection('users').authWithOAuth2({ provider: 'google' });

        // Check if authentication was successful
        if (pb.authStore.isValid) {
            console.log('Logged in with token:', pb.authStore.token);
            console.log('User ID:', pb.authStore.model.id);
        } else {
            console.log('Login failed');
        }
    } catch (error) {
        console.error('Login error:', error);
    }
});

