import { initializeApp } from 'firebase/app'
import { getAuth, type User } from 'firebase/auth'
import { GoogleAuthProvider, signInWithPopup } from 'firebase/auth'

const config = {
  apiKey: process.env.VITE_GCP_API_KEY,
  authDomain: process.env.VITE_GCP_AUTH_DOMAIN,
}

export const firebaseApp = initializeApp(config)
export const firebaseAuth = getAuth(firebaseApp)

export const signInWithGooglePopup = async () => {
  const provider = new GoogleAuthProvider()

  return await signInWithPopup(firebaseAuth, provider)
}

export const getCurrentUser = (): Promise<User | null> => {
  return new Promise((resolve) => {
    const unsubscribe = firebaseAuth.onAuthStateChanged(
      (user) => {
      unsubscribe();
      resolve(user);
    });
  });
}
