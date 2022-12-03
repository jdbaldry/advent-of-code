{ pkgs ? import <nixpkgs> }:

with pkgs;
mkShell {
  buildInputs = [
    graphviz
    gnumake
    go-jsonnet
    golangci-lint
    go_1_19
    gopls
    gotools
    jsonnet-tool
    jq
    mitscheme
    rlwrap
  ];
  shellHook = ''
    export PATH=~/bin/:$PATH
  '';
}
