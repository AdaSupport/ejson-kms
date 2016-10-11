# 3.0.0 - October 11th, 2016

Removed the `export` keyword from BASH formatter.
This is a breaking change that provides more secure defaults. Specifically,
you will need to export the environment variables yourself if your app needs
them outside of the BASH script.

# 2.0.0 - October 10th, 2016

Changed BASH escaping in export, now uses single quotes and no string processing.
This is a breaking change, but is necessary to preserve multi-line strings such
as TLS keys when using the `eval "$(ejson-kms export)"` idiom.

# 1.0.1 - October 6th, 2016

Fixed `echo "foo\nbar" | ejson-kms add` previously added only the first line

# 1.0.0 - September 28th, 2016

Initial release of `ejson-kms`
