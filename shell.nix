{ pkgs ? import <nixpkgs> }:

with pkgs;
mkShell {
  buildInputs = [
    gnumake
    go-jsonnet
    go_1_17
    gopls
    jsonnet-tool
    jq
    mitscheme
    rlwrap
  ];
  shellHook = ''
    export PATH=~/bin/:$PATH
  '';
}
