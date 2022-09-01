provider "mailman3" {
  base_url   = "https://mailman.example.org"              # optionally use MAILMAN3_BASE_URL env var
  username   = "restapi"                                  # optionally use MAILMAN3_USERNAME env var
  secret_key = "rmtuS1Uj1bIC08QFYGW18GfSHAbkPqdsuYynNudw" # optionally use MAILMAN3_PASSWORD env var
}
