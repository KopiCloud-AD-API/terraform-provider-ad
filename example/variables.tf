variable "token" {
    type        = string
    description = "Either a Bearer or Basic Authentication Token"
}

variable "oforero11_password" {
    type        = string
    description = "The password for the user"
    sensitive   = true
}