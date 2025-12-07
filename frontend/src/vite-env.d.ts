interface ImportMetaEnv {
  readonly APP_PORT: number,
  readonly VITE_GCP_API_KEY: string,
  readonly VITE_GCP_AUTH_DOMAIN: string,
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
