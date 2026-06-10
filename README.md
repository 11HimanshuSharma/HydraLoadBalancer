                Load Balancer
                       |
          ┌────────────┼────────────┐
          │            │            │

      Backend      Backend      Backend

       Alive        Alive         Dead

       Conn=7       Conn=2       Conn=0

       Fail=0       Fail=1       Fail=8