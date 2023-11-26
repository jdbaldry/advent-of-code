{ pkgs ? import <nixpkgs> }:

with pkgs;
mkShell {
  buildInputs = [
    # Go
    delve
    gdb
    go_1_19
    gofumpt
    golangci-lint
    gopls
    gotools

    # Jsonnet
    go-jsonnet
    jsonnet-tool

    # Scheme
    mitscheme

    # Tools
    gnumake
    graphviz
    jq
    rlwrap
  ];
  hardeningDisable = [ "fortify" ];
}
