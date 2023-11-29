(
    cd ../contracts
    for fn in "Aggregator" "IUniswapV2Factory" "IUniswapV2Pair"
    do
        echo ${fn}
        solc --abi ${fn}.sol -o abis --overwrite >/dev/null 2>&1
        solc --bin ${fn}.sol -o bins --overwrite >/dev/null 2>&1
        mkdir -p ../pkg/abis/${fn}
        lc=$(echo "$fn" | awk '{print tolower($0)}')
        abigen --bin=bins/${fn}.bin --abi=abis/${fn}.abi --pkg=${lc} --out=../pkg/abis/${fn}/${fn}.go
    done
    rm -r abis/* bins/*
    touch abis/.gitkeep 
    touch bins/.gitkeep
)