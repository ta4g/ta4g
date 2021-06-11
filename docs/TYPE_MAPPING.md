# Type Mappings

The ta4j library has a lot of types, for example `org.jfree.data.time.TimeSeries`, these types are mapped to Golang interfaces or structs where appropriate.

These types have been re-written from scratch to have a golang style interface.  
Looking closely you will notice they only loosely represent the original interface and no Java code has been copied into this repo.

| Java Type | Golang Type |
| --- | --- |
| org.jfree.data.time.TimeSeries | data/time/time_series/TimeSeries |
| org.ta4j.core.Bar | data/interval/bar/Bar |