{ pkgs, lib, config, inputs, ... }:

{
  # Enable reading of .env for secrets
  dotenv.enable = true;

  # https://devenv.sh/packages/
  packages = [
    pkgs.git
    pkgs.go-migrate
    pkgs.hey
  ];

  # https://devenv.sh/languages/
  languages.go.enable = true;
  languages.typescript.enable = true;
  languages.javascript.enable = true;
  languages.javascript.npm.enable = true;
  languages.javascript.npm.package = pkgs.nodejs_21;
  languages.javascript.npm.install.enable = true;

  # postgres
  services.postgres.enable = true;
  services.postgres.initialScript = ''
    CREATE DATABASE metrix;
    CREATE role metrix WITH LOGIN PASSWORD 'pass';
    GRANT ALL ON DATABASE metrix TO metrix;
    ALTER DATABASE metrix OWNER to metrix;
  '';

  services.postgres.listen_addresses = "localhost";
  services.postgres.port = 5432;

  env.METRIX_DB_DSN = "postgres://metrix:pass@localhost:5432/metrix?sslmode=disable";

  # https://devenv.sh/scripts/
  scripts.hello.exec = ''
    echo hello from $GREET
  '';

  enterShell = ''
    hello
    git --version
  '';

  # https://devenv.sh/tasks/
  # tasks = {
  #   "myproj:setup".exec = "mytool build";
  #   "devenv:enterShell".after = [ "myproj:setup" ];
  # };

  # https://devenv.sh/tests/
  enterTest = ''
    echo "Running tests"
    git --version | grep --color=auto "${pkgs.git.version}"
  '';

  # https://devenv.sh/pre-commit-hooks/
  # pre-commit.hooks.shellcheck.enable = true;

  # See full reference at https://devenv.sh/reference/options/
}
