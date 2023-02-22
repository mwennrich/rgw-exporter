# rgw-exporter

- queries a ceph-endpoint for ceph-usage
- needs a credentials (access_key, secret_key) with admin privileges (e.g. dashboard user)

## sample metrics output (QUERY_ENTRIES=true)

```text
# HELP rgw_bucket_bytes Shows rgw bucket bytes
# TYPE rgw_bucket_bytes gauge
rgw_bucket_bytes{bucket="test1215",s3_endpoint="https://s3.example.com",user="user-a"} 3.749720064e+09
rgw_bucket_bytes{bucket="test2483",s3_endpoint="https://s3.example.com",user="user-a"} 8.05310464e+08

# HELP rgw_bucket_bytes_received Shows rgw bucket received traffic in Bytes
# TYPE rgw_bucket_bytes_received counter
rgw_bucket_bytes_received{bucket="",category="list_buckets",owner="user-a",s3_endpoint="https://s3.example.com",user="user-a"} 0
rgw_bucket_bytes_received{bucket="-",category="get_bucket_object_lock",owner="user-a",s3_endpoint="https://s3.example.com",user="user-a"} 0
rgw_bucket_bytes_received{bucket="-",category="get_obj",owner="user-a",s3_endpoint="https://s3.example.com",user="user-a"} 0
rgw_bucket_bytes_received{bucket="test2483",category="get_bucket_location",owner="user-a",s3_endpoint="https://s3.example.com",user="user-a"} 0
rgw_bucket_bytes_received{bucket="test2483",category="list_bucket",owner="user-a",s3_endpoint="https://s3.example.com",user="user-a"} 0
rgw_bucket_bytes_received{bucket="test2483",category="put_obj",owner="user-a",s3_endpoint="https://s3.example.com",user="user-a"} 215

# HELP rgw_bucket_bytes_sent Shows rgw bucket sent traffic in Bytes
# TYPE rgw_bucket_bytes_sent counter
rgw_bucket_bytes_sent{bucket="",category="list_buckets",owner="user-a",s3_endpoint="https://s3.example.com",user="user-a"} 3582
rgw_bucket_bytes_sent{bucket="-",category="get_bucket_object_lock",owner="user-a",s3_endpoint="https://s3.example.com",user="user-a"} 233
rgw_bucket_bytes_sent{bucket="-",category="get_obj",owner="user-a",s3_endpoint="https://s3.example.com",user="user-a"} 0
rgw_bucket_bytes_sent{bucket="test2483",category="get_bucket_location",owner="user-a",s3_endpoint="https://s3.example.com",user="user-a"} 258
rgw_bucket_bytes_sent{bucket="test2483",category="list_bucket",owner="user-a",s3_endpoint="https://s3.example.com",user="user-a"} 979
rgw_bucket_bytes_sent{bucket="test2483",category="put_obj",owner="user-a",s3_endpoint="https://s3.example.com",user="user-a"} 0

# HELP rgw_bucket_objects Shows rgw bucket number of objects
# TYPE rgw_bucket_objects gauge
rgw_bucket_objects{bucket="test1215",s3_endpoint="https://s3.example.com",user="user-a"} 40
rgw_bucket_objects{bucket="test2483",s3_endpoint="https://s3.example.com",user="user-a"} 49

# HELP rgw_bucket_ops Shows rgw bucket ops
# TYPE rgw_bucket_ops counter
rgw_bucket_ops{bucket="",category="list_buckets",owner="user-a",s3_endpoint="https://s3.example.com",user="user-a"} 6
rgw_bucket_ops{bucket="-",category="get_bucket_object_lock",owner="user-a",s3_endpoint="https://s3.example.com",user="user-a"} 1
rgw_bucket_ops{bucket="-",category="get_obj",owner="user-a",s3_endpoint="https://s3.example.com",user="user-a"} 2
rgw_bucket_ops{bucket="test2483",category="get_bucket_location",owner="user-a",s3_endpoint="https://s3.example.com",user="user-a"} 2
rgw_bucket_ops{bucket="test2483",category="list_bucket",owner="user-a",s3_endpoint="https://s3.example.com",user="user-a"} 3
rgw_bucket_ops{bucket="test2483",category="put_obj",owner="user-a",s3_endpoint="https://s3.example.com",user="user-a"} 1

# HELP rgw_bucket_ops_successful Shows rgw bucket successful ops
# TYPE rgw_bucket_ops_successful counter
rgw_bucket_ops_successful{bucket="",category="list_buckets",owner="user-a",s3_endpoint="https://s3.example.com",user="user-a"} 6
rgw_bucket_ops_successful{bucket="-",category="get_bucket_object_lock",owner="user-a",s3_endpoint="https://s3.example.com",user="user-a"} 0
rgw_bucket_ops_successful{bucket="-",category="get_obj",owner="user-a",s3_endpoint="https://s3.example.com",user="user-a"} 0
rgw_bucket_ops_successful{bucket="test2483",category="get_bucket_location",owner="user-a",s3_endpoint="https://s3.example.com",user="user-a"} 2
rgw_bucket_ops_successful{bucket="test2483",category="list_bucket",owner="user-a",s3_endpoint="https://s3.example.com",user="user-a"} 3
rgw_bucket_ops_successful{bucket="test2483",category="put_obj",owner="user-a",s3_endpoint="https://s3.example.com",user="user-a"} 1

# HELP rgw_category_bytes_received Shows rgw category received traffic in Bytes
# TYPE rgw_category_bytes_received counter
rgw_category_bytes_received{category="get_bucket_location",s3_endpoint="https://s3.example.com",user="user-a"} 0
rgw_category_bytes_received{category="get_bucket_object_lock",s3_endpoint="https://s3.example.com",user="user-a"} 0
rgw_category_bytes_received{category="get_obj",s3_endpoint="https://s3.example.com",user="user-a"} 0
rgw_category_bytes_received{category="list_bucket",s3_endpoint="https://s3.example.com",user="user-a"} 0
rgw_category_bytes_received{category="list_buckets",s3_endpoint="https://s3.example.com",user="user-a"} 0
rgw_category_bytes_received{category="put_obj",s3_endpoint="https://s3.example.com",user="user-a"} 215

# HELP rgw_category_bytes_sent Shows rgw category sent traffic in Bytes
# TYPE rgw_category_bytes_sent counter
rgw_category_bytes_sent{category="get_bucket_location",s3_endpoint="https://s3.example.com",user="user-a"} 258
rgw_category_bytes_sent{category="get_bucket_object_lock",s3_endpoint="https://s3.example.com",user="user-a"} 233
rgw_category_bytes_sent{category="get_obj",s3_endpoint="https://s3.example.com",user="user-a"} 0
rgw_category_bytes_sent{category="list_bucket",s3_endpoint="https://s3.example.com",user="user-a"} 979
rgw_category_bytes_sent{category="list_buckets",s3_endpoint="https://s3.example.com",user="user-a"} 3582
rgw_category_bytes_sent{category="put_obj",s3_endpoint="https://s3.example.com",user="user-a"} 0

# HELP rgw_category_ops Shows rgw category ops
# TYPE rgw_category_ops counter
rgw_category_ops{category="get_bucket_location",s3_endpoint="https://s3.example.com",user="user-a"} 2
rgw_category_ops{category="get_bucket_object_lock",s3_endpoint="https://s3.example.com",user="user-a"} 1
rgw_category_ops{category="get_obj",s3_endpoint="https://s3.example.com",user="user-a"} 2
rgw_category_ops{category="list_bucket",s3_endpoint="https://s3.example.com",user="user-a"} 3
rgw_category_ops{category="list_buckets",s3_endpoint="https://s3.example.com",user="user-a"} 6
rgw_category_ops{category="put_obj",s3_endpoint="https://s3.example.com",user="user-a"} 1

# HELP rgw_category_ops_successful Shows rgw category successful ops
# TYPE rgw_category_ops_successful counter
rgw_category_ops_successful{category="get_bucket_location",s3_endpoint="https://s3.example.com",user="user-a"} 2
rgw_category_ops_successful{category="get_bucket_object_lock",s3_endpoint="https://s3.example.com",user="user-a"} 0
rgw_category_ops_successful{category="get_obj",s3_endpoint="https://s3.example.com",user="user-a"} 0
rgw_category_ops_successful{category="list_bucket",s3_endpoint="https://s3.example.com",user="user-a"} 3
rgw_category_ops_successful{category="list_buckets",s3_endpoint="https://s3.example.com",user="user-a"} 6
rgw_category_ops_successful{category="put_obj",s3_endpoint="https://s3.example.com",user="user-a"} 1

# HELP rgw_total_bytes Shows rgw total bytes
# TYPE rgw_total_bytes gauge
rgw_total_bytes{s3_endpoint="https://s3.example.com",user="user-a"} 4.555030528e+09

# HELP rgw_total_bytes_received Shows rgw total received traffic in Bytes
# TYPE rgw_total_bytes_received counter
rgw_total_bytes_received{s3_endpoint="https://s3.example.com",user="user-a"} 215

# HELP rgw_total_bytes_sent Shows rgw total sent traffic in Bytes
# TYPE rgw_total_bytes_sent counter
rgw_total_bytes_sent{s3_endpoint="https://s3.example.com",user="user-a"} 5052

# HELP rgw_total_objects Shows rgw total number of objects
# TYPE rgw_total_objects gauge
rgw_total_objects{s3_endpoint="https://s3.example.com",user="user-a"} 89

# HELP rgw_total_ops Shows rgw total ops
# TYPE rgw_total_ops counter
rgw_total_ops{s3_endpoint="https://s3.example.com",user="user-a"} 15

# HELP rgw_total_ops_successful Shows rgw total successful ops
# TYPE rgw_total_ops_successful counter
rgw_total_ops_successful{s3_endpoint="https://s3.example.com",user="user-a"} 12
```
