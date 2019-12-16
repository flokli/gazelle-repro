# gazelle visibility repro

```
bazel run //:gazelle
bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro go_repositories.bzl%go_repositories -prune
bazel build //:someproject.com
```

Bug:
```
[nix-shell:~/dev/foo/go-repro]$ bazel run //:gazelle
INFO: Analyzed target //:gazelle (27 packages loaded, 131 targets configured).
INFO: Found 1 target...
Target //:gazelle up-to-date:
  bazel-bin/gazelle-runner.bash
  bazel-bin/gazelle
INFO: Elapsed time: 0.586s, Critical Path: 0.01s
INFO: 0 processes.
INFO: Build completed successfully, 1 total action
INFO: Build completed successfully, 1 total action

[nix-shell:~/dev/foo/go-repro]$ bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro go_repositories.bzl%go_repositories -prune
INFO: Analyzed target //:gazelle (1 packages loaded, 2 targets configured).
INFO: Found 1 target...
Target //:gazelle up-to-date:
  bazel-bin/gazelle-runner.bash
  bazel-bin/gazelle
INFO: Elapsed time: 0.179s, Critical Path: 0.01s
INFO: 0 processes.
INFO: Build completed successfully, 1 total action
INFO: Running command line: bazel-bin/gazelle -bazel_run update-repos '-from_file=go.mod' -to_macro go_INFO: Build completed successfully, 1 total action

[nix-shell:~/dev/foo/go-repro]$ bazel build //:someproject.com
ERROR: /home/flokli/.cache/bazel/_bazel_flokli/71894a93d52ccbda3bc58b835e92a444/external/com_google_cloud_go_storage/BUILD.bazel:3:1: in go_library rule @com_google_cloud_go_storage//:go_default_library: target '@com_google_cloud_go//internal:go_default_library' is not visible from target '@com_google_cloud_go_storage//:go_default_library'. Check the visibility declaration of the former target if you think the dependency is legitimate
ERROR: /home/flokli/.cache/bazel/_bazel_flokli/71894a93d52ccbda3bc58b835e92a444/external/com_google_cloud_go_storage/BUILD.bazel:3:1: in go_library rule @com_google_cloud_go_storage//:go_default_library: target '@com_google_cloud_go//internal/optional:go_default_library' is not visible from target '@com_google_cloud_go_storage//:go_default_library'. Check the visibility declaration of the former target if you think the dependency is legitimate
ERROR: /home/flokli/.cache/bazel/_bazel_flokli/71894a93d52ccbda3bc58b835e92a444/external/com_google_cloud_go_storage/BUILD.bazel:3:1: in go_library rule @com_google_cloud_go_storage//:go_default_library: target '@com_google_cloud_go//internal/trace:go_default_library' is not visible from target '@com_google_cloud_go_storage//:go_default_library'. Check the visibility declaration of the former target if you think the dependency is legitimate
ERROR: /home/flokli/.cache/bazel/_bazel_flokli/71894a93d52ccbda3bc58b835e92a444/external/com_google_cloud_go_storage/BUILD.bazel:3:1: in go_library rule @com_google_cloud_go_storage//:go_default_library: target '@com_google_cloud_go//internal/version:go_default_library' is not visible from target '@com_google_cloud_go_storage//:go_default_library'. Check the visibility declaration of the former target if you think the dependency is legitimate
ERROR: Analysis of target '//:someproject.com' failed; build aborted: Analysis of target '@com_google_cloud_go_storage//:go_default_library' failed; build aborted
INFO: Elapsed time: 0.716s
INFO: 0 processes.
FAILED: Build did NOT complete successfully (104 packages loaded, 1171 targets configured)

[nix-shell:~/dev/foo/go-repro]$ bazel version
Build label: 1.2.1- (@non-git)
Build target: bazel-out/k8-opt/bin/src/main/java/com/google/devtools/build/lib/bazel/BazelServer_deploy.jar
Build time: Tue Jan 1 00:00:00 1980 (315532800)
Build timestamp: 315532800
Build timestamp as int: 315532800
```
