# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Farmhub < Formula
  desc ""
  homepage "https://farmhub.ag"
  version "1.5.19"

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/thefarmhub/farmhub-cli/releases/download/v1.5.19/farmhub_Darwin_x86_64.tar.gz"
      sha256 "cf68a7eb672a5c16edce4525f7f705b7a17508c2cad31500a5d5754736a83ea0"

      def install
        bin.install "farmhub"
      end
    end
    if Hardware::CPU.arm?
      url "https://github.com/thefarmhub/farmhub-cli/releases/download/v1.5.19/farmhub_Darwin_arm64.tar.gz"
      sha256 "30d313bc5aa18d73cb19d77d6bd728ed87fa9126950db6eb312985b0622daf8b"

      def install
        bin.install "farmhub"
      end
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      url "https://github.com/thefarmhub/farmhub-cli/releases/download/v1.5.19/farmhub_Linux_x86_64.tar.gz"
      sha256 "1baafbc5417dd7b69a558413d4b06d348e5f7cc666db37743dc386c3b17c312d"

      def install
        bin.install "farmhub"
      end
    end
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/thefarmhub/farmhub-cli/releases/download/v1.5.19/farmhub_Linux_arm64.tar.gz"
      sha256 "75e61bab4565141a83aca90402f00ee5c012ee7d22cabad8f1354afa44d9d7bc"

      def install
        bin.install "farmhub"
      end
    end
  end

  test do
    system "#{bin}/farmhub version"
  end
end
