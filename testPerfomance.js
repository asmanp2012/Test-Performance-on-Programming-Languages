const performanceTest = () => {
    "use strict";
    const delta = () => 10 * 15 / 20  * 40;

    let duration = 0;
    let count = 0;
    const step = 100_000;
    const StartTime = Date.now();
    
    for (; duration < 1000; count += step) {
        for (let index = 0; index < step; index++) {
            delta();
        }
        duration = Date.now() - StartTime;
    }
    
    console.log("count = %s", count);
};

setTimeout(performanceTest, 1000);