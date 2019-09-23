package main

import "github.com/picfight/coin_knife/fileproc"

func PicfightCoinFileNameGenerator(data string) string {
	data = fileproc.Replace(data, "btcsuite", "btcziggurat")
	return data
}

func PicfightCoinFileGenerator(data string) string {
	data = fileproc.Replace(data, "btcsuite/btcd", "btcziggurat/btcd")
	data = fileproc.Replace(data, "btcsuite\btcd", "btcziggurat\btcd")

	data = fileproc.Replace(data, "btcsuite/btcutil", "btcziggurat/btcutil")
	data = fileproc.Replace(data, "btcsuite\btcutil", "btcziggurat\btcutil")

	data = fileproc.Replace(data, "btcsuite/btcwallet", "btcziggurat/btcwallet")
	data = fileproc.Replace(data, "btcsuite\btcwallet", "btcziggurat\btcwallet")

	return data
}
