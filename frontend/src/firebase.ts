import { initializeApp } from 'firebase/app'
import { getAuth } from 'firebase/auth'
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
