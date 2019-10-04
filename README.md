# boilr-metrics

This is a [boilr template](https://github.com/tmrts/boilr) for creating a custom metrics provider. Create a custom metrics provider to expose additional metrics from the Drone database. Get started by cloning the project and installing the template:

```console
$ cd boilr-metrics
$ boilr template save . drone-metrics -f
```

create a project in directory my-metrics:

```console
$ boilr template use drone-metrics my-metrics
```

enter the go module name:

```text
[?] Please choose a value for "GoModule" [default: "github.com/foo/bar":
```

enter the docker registry name for this project:

```text
[?] Please choose a value for "DockerRepository" [default: "foo/bar"]:
```
