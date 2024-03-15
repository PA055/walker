{
  mkShell,
  mkGoEnv,
  gomod2nix,
  go,
  gtk4-layer-shell,
}:

mkShell {
  packages = [
    (mkGoEnv {pwd = ./.;})

    # Build
    gomod2nix
    go

    # Dependencies
    gtk4-layer-shell
  ];
}
