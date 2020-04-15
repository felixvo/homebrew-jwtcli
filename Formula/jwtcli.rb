# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
# PLEASE REMOVE ALL GENERATED COMMENTS BEFORE SUBMITTING YOUR PULL REQUEST!
class Jwtcli < Formula
  desc "JWT CLI tool"
  homepage "https://github.com/felixvo/jwtcli"
  url "https://github.com/felixvo/jwtcli/raw/master/jwtcli.tar.gz"
  version "1.0.0"
  sha256 "8a3d2be7e21b8715854dcd557c35acbdb09479844a212644fdf2e25b38bde47e"

  # depends_on "cmake" => :build

  def install
    # ENV.deparallelize  # if your formula fails when building in parallel
    # Remove unrecognized options if warned by configure
    #system "ln -s jwtcli /usr/bin/jwtcli"
    # system "cmake", ".", *std_cmake_args
    bin.install "jwtcli"
  end

  test do
    # `test do` will create, run in and delete a temporary directory.
    #
    # This test will fail and we won't accept that! For Homebrew/homebrew-core
    # this will need to be a test that verifies the functionality of the
    # software. Run the test with `brew test jwtcli`. Options passed
    # to `brew install` such as `--HEAD` also need to be provided to `brew test`.
    #
    # The installed folder is not in the path, so use the entire path to any
    # executables being tested: `system "#{bin}/program", "do", "something"`.
    system "false"
  end
end
