# rgw-exporter

- queries a ceph-endpoint for ceph-usage
- needs a credentials (access_key, secret_key) with admin privileges (e.g. dashboard user)

## sample metrics output (QUERY_ENTRIES=true)

```text
# HELP rgw_bucket_bytes_received Shows rgw bucket received traffic in Bytes
# TYPE rgw_bucket_bytes_received counter
rgw_bucket_bytes_received{bucket="",category="list_buckets",owner="anonymous",s3_endpoint="https://s3.site-a.example.com",user="anonymous"} 0
rgw_bucket_bytes_received{bucket="",category="list_buckets",owner="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 0
rgw_bucket_bytes_received{bucket="rook-ceph-bucket-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",category="create_bucket",owner="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 0
rgw_bucket_bytes_received{bucket="rook-ceph-bucket-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",category="delete_obj",owner="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 0
rgw_bucket_bytes_received{bucket="rook-ceph-bucket-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",category="get_obj",owner="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 0
rgw_bucket_bytes_received{bucket="rook-ceph-bucket-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",category="put_obj",owner="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 11886
rgw_bucket_bytes_received{bucket="testbucket-site-a",category="create_bucket",owner="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 0
rgw_bucket_bytes_received{bucket="testbucket-site-a",category="delete_obj",owner="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 0
rgw_bucket_bytes_received{bucket="testbucket-site-a",category="get_obj",owner="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 0
rgw_bucket_bytes_received{bucket="testbucket-site-a",category="put_obj",owner="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 4352

# HELP rgw_bucket_bytes_sent Shows rgw bucket sent traffic in Bytes
# TYPE rgw_bucket_bytes_sent counter
rgw_bucket_bytes_sent{bucket="",category="list_buckets",owner="anonymous",s3_endpoint="https://s3.site-a.example.com",user="anonymous"} 231
rgw_bucket_bytes_sent{bucket="",category="list_buckets",owner="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 29871
rgw_bucket_bytes_sent{bucket="rook-ceph-bucket-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",category="create_bucket",owner="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 0
rgw_bucket_bytes_sent{bucket="rook-ceph-bucket-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",category="delete_obj",owner="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 0
rgw_bucket_bytes_sent{bucket="rook-ceph-bucket-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",category="get_obj",owner="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 11886
rgw_bucket_bytes_sent{bucket="rook-ceph-bucket-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",category="put_obj",owner="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 0
rgw_bucket_bytes_sent{bucket="testbucket-site-a",category="create_bucket",owner="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 0
rgw_bucket_bytes_sent{bucket="testbucket-site-a",category="delete_obj",owner="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 0
rgw_bucket_bytes_sent{bucket="testbucket-site-a",category="get_obj",owner="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 4352
rgw_bucket_bytes_sent{bucket="testbucket-site-a",category="put_obj",owner="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 0

# HELP rgw_bucket_ops Shows rgw bucket ops
# TYPE rgw_bucket_ops counter
rgw_bucket_ops{bucket="",category="list_buckets",owner="anonymous",s3_endpoint="https://s3.site-a.example.com",user="anonymous"} 1
rgw_bucket_ops{bucket="",category="list_buckets",owner="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 68
rgw_bucket_ops{bucket="rook-ceph-bucket-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",category="create_bucket",owner="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 566
rgw_bucket_ops{bucket="rook-ceph-bucket-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",category="delete_obj",owner="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 1132
rgw_bucket_ops{bucket="rook-ceph-bucket-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",category="get_obj",owner="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 566
rgw_bucket_ops{bucket="rook-ceph-bucket-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",category="put_obj",owner="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 566
rgw_bucket_ops{bucket="testbucket-site-a",category="create_bucket",owner="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 1
rgw_bucket_ops{bucket="testbucket-site-a",category="delete_obj",owner="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 68
rgw_bucket_ops{bucket="testbucket-site-a",category="get_obj",owner="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 204
rgw_bucket_ops{bucket="testbucket-site-a",category="put_obj",owner="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 68

# HELP rgw_bucket_ops_successful Shows rgw bucket sucessfull ops
# TYPE rgw_bucket_ops_successful counter
rgw_bucket_ops_successful{bucket="",category="list_buckets",owner="anonymous",s3_endpoint="https://s3.site-a.example.com",user="anonymous"} 1
rgw_bucket_ops_successful{bucket="",category="list_buckets",owner="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 68
rgw_bucket_ops_successful{bucket="rook-ceph-bucket-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",category="create_bucket",owner="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 566
rgw_bucket_ops_successful{bucket="rook-ceph-bucket-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",category="delete_obj",owner="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 1132
rgw_bucket_ops_successful{bucket="rook-ceph-bucket-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",category="get_obj",owner="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 566
rgw_bucket_ops_successful{bucket="rook-ceph-bucket-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",category="put_obj",owner="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 566
rgw_bucket_ops_successful{bucket="testbucket-site-a",category="create_bucket",owner="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 1
rgw_bucket_ops_successful{bucket="testbucket-site-a",category="delete_obj",owner="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 68
rgw_bucket_ops_successful{bucket="testbucket-site-a",category="get_obj",owner="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 204
rgw_bucket_ops_successful{bucket="testbucket-site-a",category="put_obj",owner="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 68

# HELP rgw_category_bytes_received Shows rgw category received traffic in Bytes
# TYPE rgw_category_bytes_received counter
rgw_category_bytes_received{category="create_bucket",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 0
rgw_category_bytes_received{category="create_bucket",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 0
rgw_category_bytes_received{category="delete_obj",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 0
rgw_category_bytes_received{category="delete_obj",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 0
rgw_category_bytes_received{category="get_obj",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 0
rgw_category_bytes_received{category="get_obj",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 0
rgw_category_bytes_received{category="list_buckets",s3_endpoint="https://s3.site-a.example.com",user="anonymous"} 0
rgw_category_bytes_received{category="list_buckets",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 0
rgw_category_bytes_received{category="put_obj",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 4352
rgw_category_bytes_received{category="put_obj",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 11886

# HELP rgw_category_bytes_sent Shows rgw category sent traffic in Bytes
# TYPE rgw_category_bytes_sent counter
rgw_category_bytes_sent{category="create_bucket",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 0
rgw_category_bytes_sent{category="create_bucket",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 0
rgw_category_bytes_sent{category="delete_obj",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 0
rgw_category_bytes_sent{category="delete_obj",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 0
rgw_category_bytes_sent{category="get_obj",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 4352
rgw_category_bytes_sent{category="get_obj",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 11886
rgw_category_bytes_sent{category="list_buckets",s3_endpoint="https://s3.site-a.example.com",user="anonymous"} 231
rgw_category_bytes_sent{category="list_buckets",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 29871
rgw_category_bytes_sent{category="put_obj",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 0
rgw_category_bytes_sent{category="put_obj",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 0

# HELP rgw_category_ops Shows rgw category ops
# TYPE rgw_category_ops counter
rgw_category_ops{category="create_bucket",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 1
rgw_category_ops{category="create_bucket",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 566
rgw_category_ops{category="delete_obj",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 68
rgw_category_ops{category="delete_obj",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 1132
rgw_category_ops{category="get_obj",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 204
rgw_category_ops{category="get_obj",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 566
rgw_category_ops{category="list_buckets",s3_endpoint="https://s3.site-a.example.com",user="anonymous"} 1
rgw_category_ops{category="list_buckets",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 68
rgw_category_ops{category="put_obj",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 68
rgw_category_ops{category="put_obj",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 566

# HELP rgw_category_ops_successful Shows rgw category sucessfull ops
# TYPE rgw_category_ops_successful counter
rgw_category_ops_successful{category="create_bucket",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 1
rgw_category_ops_successful{category="create_bucket",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 566
rgw_category_ops_successful{category="delete_obj",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 68
rgw_category_ops_successful{category="delete_obj",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 1132
rgw_category_ops_successful{category="get_obj",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 204
rgw_category_ops_successful{category="get_obj",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 566
rgw_category_ops_successful{category="list_buckets",s3_endpoint="https://s3.site-a.example.com",user="anonymous"} 1
rgw_category_ops_successful{category="list_buckets",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 68
rgw_category_ops_successful{category="put_obj",s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 68
rgw_category_ops_successful{category="put_obj",s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 566

# HELP rgw_total_bytes_received Shows rgw total received traffic in Bytes
# TYPE rgw_total_bytes_received counter
rgw_total_bytes_received{s3_endpoint="https://s3.site-a.example.com",user="anonymous"} 0
rgw_total_bytes_received{s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 4352
rgw_total_bytes_received{s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 11886

# HELP rgw_total_bytes_sent Shows rgw total sent traffic in Bytes
# TYPE rgw_total_bytes_sent counter
rgw_total_bytes_sent{s3_endpoint="https://s3.site-a.example.com",user="anonymous"} 231
rgw_total_bytes_sent{s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 34223
rgw_total_bytes_sent{s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 11886

# HELP rgw_total_ops Shows rgw total ops
# TYPE rgw_total_ops counter
rgw_total_ops{s3_endpoint="https://s3.site-a.example.com",user="anonymous"} 1
rgw_total_ops{s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 409
rgw_total_ops{s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 2830

# HELP rgw_total_ops_successful Shows rgw total sucessfull ops
# TYPE rgw_total_ops_successful counter
rgw_total_ops_successful{s3_endpoint="https://s3.site-a.example.com",user="anonymous"} 1
rgw_total_ops_successful{s3_endpoint="https://s3.site-a.example.com",user="tenant$5820c4e7-fbd4-4e4b-a40b-2b83eb34bbe1_s3prober"} 409
rgw_total_ops_successful{s3_endpoint="https://s3.site-a.example.com",user="rook-ceph-internal-s3-user-checker-ef258790-dbe2-4b5f-bcab-8dc36dd33801"} 2830
```
