# Contributing

By participating to this project, you agree to abide our [code of
conduct](/CODE_OF_CONDUCT.md).

## Setup your environment

`xela` is written in [Go](https://golang.org/).

Prerequisites:

* [Go 1.11+](https://golang.org/doc/install)

Clone `xela` from source into `$GOPATH`:

```sh
$ mkdir -p $GOPATH/src/github.com/mattstratton/xela
$ cd $_
$ git clone git@github.com:mattstratton/xela.git .
```

### Git remote setup

Change our remote to be named `upstream`:

```sh
$ git remote rename origin upstream
```

Add your fork as `origin`:

```sh
$ git remote add fork git@github.com:you/xela.git
```

## Making changes

### Testing changes

TODO add testing stuff here

### Create a commit

Commit messages should be well formatted.
Start your commit message with a title in the imperative, i.e., "Updates function foo" vs "Updated function foo". Capitalize it.

The title must be followed with a newline, then a more detailed description.

Please reference any GitHub issues on the last line of the commit message (e.g. `See #123`, `Closes #123`, `Fixes #123`).

An example:

```
Add example for --debug flag

I added an example to the docs of the `--debug` flag to make
the usage more clear.

Fixes #284
```

### Branching and Pull Requests

(inspired by [Katrina Owen](kytrinyx)'s [excellent blog post](https://splice.com/blog/contributing-open-source-git-repositories-go/))

Assuming that the `you/xela` repo is at `origin`, and `mattstratton/xela` is at `upstream`, here's how to create a pull request:

```sh

$ git checkout -b make-thing-awesome
$ git commit -a myfile.go
$ git commit -s -m "Make thing more awesome"
$ git push origin make-thing-awesome

```

Don't forget to keep up to date with `upstream`:

```sh
$ git fetch upstream
$ git reset --hard upstream/master
```

## Developer Certification of Origin (DCO)

Licensing is very important to open source projects. It helps ensure the software continues to be available under the terms that the author desired.

This project uses [the MIT license](https://github.com/mattstratton/xela/blob/master/LICENSE).

The license tells you what rights you have that are provided by the copyright holder. It is important that the contributor fully understands what rights they are licensing and agrees to them. Sometimes the copyright holder isn't the contributor, such as when the contributor is doing work on behalf of a company.

To make a good faith effort to ensure these criteria are met, we requires the Developer Certificate of Origin (DCO) process to be followed.

The DCO is an attestation attached to every contribution made by every developer. In the commit message of the contribution, the developer simply adds a Signed-off-by statement and thereby agrees to the DCO, which you can find below or at <http://developercertificate.org/>.

```
Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the
    best of my knowledge, is covered under an appropriate open
    source license and I have the right under that license to   
    submit that work with modifications, whether created in whole
    or in part by me, under the same open source license (unless
    I am permitted to submit under a different license), as
    Indicated in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including
    all personal information I submit with it, including my
    sign-off) is maintained indefinitely and may be redistributed
    consistent with this project or the open source license(s)
    involved.
```

#### DCO Sign-Off Methods

The DCO requires a sign-off message in the following format appear on each commit in the pull request:

```
Signed-off-by: George Bluth <george.bluth@bluthcompany.com>
```

The DCO text can either be manually added to your commit body, or you can add either **-s** or **--signoff** to your usual git commit commands. If you forget to add the sign-off you can also amend a previous commit with the sign-off by running **git commit --amend -s**. If you've pushed your changes to Github already you'll need to force push your branch after this with **git push -f**.