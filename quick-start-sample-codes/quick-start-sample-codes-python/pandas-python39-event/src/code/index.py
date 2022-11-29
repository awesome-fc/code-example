# -*- coding: utf-8 -*-
import pandas as pd

def handler(event, context):
    df = pd.DataFrame(
        {
            "A": 1.0,
            "B": pd.Timestamp("20130102"),
            "C": pd.Series(1, index=list(range(4)), dtype="float32"),
            "D": pd.Categorical(["test", "train", "test", "train"]),
            "E": "foo",
        }
    )
    print(df.head())
    return "pandas version:"+pd.__version__