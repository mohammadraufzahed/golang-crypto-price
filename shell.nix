{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
    buildInputs = [
        pkgs.podman
        pkgs.podman-compose
        pkgs.go
        pkgs.fish
    ];

    shellHook = ''
        export SHELL=$(which fish)
        exec fish
    '';
}