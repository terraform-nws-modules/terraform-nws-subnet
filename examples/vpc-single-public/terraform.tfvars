name    = ["mynet-public"]
cidr    = ["10.0.1.0/30"]
domain  = "my.local"
vpc_id  = "fd7a2090-2e34-45b8-9eca-d9d27584475d"
public  = true
has_vpc = true

acl_allowed_port_list = ["80", "31000-31004"]
