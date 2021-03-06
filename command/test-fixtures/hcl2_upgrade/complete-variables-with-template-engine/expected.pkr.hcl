
variable "env_test" {
  type    = string
  default = "${env("TEST_ENV")}"
}

locals { timestamp = regex_replace(timestamp(), "[- TZ:]", "") }

# 5 errors occurred upgrading the following block:
# unhandled "lower" call:
# there is no way to automatically upgrade the "lower" call.
# Please manually upgrade to `lower(var.example)`
# Visit https://www.packer.io/docs/templates/hcl_templates/functions/string/lower for more infos.

# unhandled "replace" call:
# there is no way to automatically upgrade the "replace" call.
# Please manually upgrade to `replace(string, substring, replacement)` or `regex_replace(string, substring, replacement)`
# Visit https://www.packer.io/docs/templates/hcl_templates/functions/string/replace or https://www.packer.io/docs/templates/hcl_templates/functions/string/regex_replace for more infos.

# unhandled "replace_all" call:
# there is no way to automatically upgrade the "replace_all" call.
# Please manually upgrade to `replace(string, substring, replacement)` or `regex_replace(string, substring, replacement)`
# Visit https://www.packer.io/docs/templates/hcl_templates/functions/string/replace or https://www.packer.io/docs/templates/hcl_templates/functions/string/regex_replace for more infos.

# unhandled "split" call:
# there is no way to automatically upgrade the "split" call.
# Please manually upgrade to `split(separator, string)`
# Visit https://www.packer.io/docs/templates/hcl_templates/functions/string/split for more infos.

# unhandled "upper" call:
# there is no way to automatically upgrade the "upper" call.
# Please manually upgrade to `upper(var.example)`
# Visit https://www.packer.io/docs/templates/hcl_templates/functions/string/upper for more infos.
locals {
  build_timestamp = "${local.timestamp}"
  iso_datetime    = "${local.timestamp}"
  lower           = "{{ lower `HELLO` }}"
  pwd             = "${path.cwd}"
  replace         = "{{ replace `b` `c` `ababa` 2 }}"
  replace_all     = "{{ replace_all `b` `c` `ababa` }}"
  split           = "{{ split `aba` `b` 1 }}"
  temp_directory  = "${path.root}"
  upper           = "{{ upper `hello` }}"
  uuid            = "${uuidv4()}"
}

source "null" "autogenerated_1" {
  communicator = "none"
}

build {
  sources = ["source.null.autogenerated_1"]

  provisioner "shell-local" {
    inline = ["echo ${local.build_timestamp}", "echo ${local.temp_directory}", "echo ${local.iso_datetime}", "echo ${local.uuid}", "echo ${var.env_test}", "echo ${local.lower}", "echo ${local.upper}", "echo ${local.pwd}", "echo ${local.replace}", "echo ${local.replace_all}", "echo ${local.split}"]
  }

}
