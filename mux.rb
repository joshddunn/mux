# https://docs.brew.sh/Cask-Cookbook

cask "mux" do
  version "0.0.1"
  sha256 "todo"

  url "https://github.com/joshddun/mux/archive/#{version}.tar.gz"
  name "mux"
  desc "Configure tmux sessions"
  homepage "https://github.com/joshddunn/mux"

  binary "mux"
end
