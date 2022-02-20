with import <nixpkgs> {};
let
  lpedit-lib = buildGoModule {
    name = "lpedit-lib";

    src = ./.;

    buildInputs = [ alsa-lib ];

    vendorSha256 = "0sjjj9z1dhilhpc8pq4154czrb79z9cm044jvn75kxcjv6v5l2m5";
  };
in mkShell {
  buildInputs = [
    lpedit-lib
  ];
}
