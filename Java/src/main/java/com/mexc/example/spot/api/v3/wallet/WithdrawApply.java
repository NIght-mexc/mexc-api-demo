package com.mexc.example.spot.api.v3.wallet;

import com.fasterxml.jackson.core.type.TypeReference;
import com.google.common.collect.ImmutableMap;
import com.google.common.collect.Maps;
import com.mexc.example.common.JsonUtil;
import com.mexc.example.common.UserDataClient;
import com.mexc.example.spot.api.v3.pojo.Withdraw;
import lombok.extern.slf4j.Slf4j;

import java.util.HashMap;
import java.util.Map;

@Slf4j
public class WithdrawApply {
    public static Withdraw withdraw(Map<String, String> params) {
        return UserDataClient.post("/api/v3/capital/withdraw", params, new TypeReference<Withdraw>() {
        });
    }

    public static void main(String[] args) {
        //withdraw apply
        HashMap<String, String> withdrawParams = Maps.newHashMap(ImmutableMap.<String, String>builder()
                .put("coin", "USDT")
                .put("address", "TLHwAkP8Ao5vHc8GR9KvwhrJHwJN142ECZ")
                .put("amount", "0.1")
                .put("netWork", "TRX")
                .put("recvWindow", "60000")
                .build());

        Object withdraw = withdraw(withdrawParams);
        log.info("===>>withdraw resp:{}", JsonUtil.toJson(withdraw));
    }
}
