# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Farmhub < Formula
  desc ""
  homepage "https://farmhub.ag"
  version "1.5.8"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/thefarmhub/farmhub-cli/releases/download/v1.5.8/farmhub_Darwin_arm64.tar.gz"
      sha256 "7de5adc7b662b9d545e787a324b7498d133512d60711edfd0017e29b68729fe6"

      def install
        bin.install "farmhub"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/thefarmhub/farmhub-cli/releases/download/v1.5.8/farmhub_Darwin_x86_64.tar.gz"
      sha256 "eee9657b3bb9846845348c873ec5d14b3f2a8750f0a039c8dac9f2fc0ce33de8"

      def install
        bin.install "farmhub"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/thefarmhub/farmhub-cli/releases/download/v1.5.8/farmhub_Linux_arm64.tar.gz"
      sha256 "5d9944c5bc1610cc6f5900f683ee051d6711013616a0f546a48101a2fbe0ae15"

      def install
        bin.install "farmhub"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/thefarmhub/farmhub-cli/releases/download/v1.5.8/farmhub_Linux_x86_64.tar.gz"
      sha256 "946f77e2f4a40abcfd7154c75f8ae7b790a4c627a92310128775823c7b460555"

      def install
        bin.install "farmhub"
      end
    end
  end

  test do
    system "#{bin}/farmhub version"
  end
end
