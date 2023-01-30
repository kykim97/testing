package testing.domain;

import java.util.*;
import lombok.*;
import testing.domain.*;
import testing.infra.AbstractEvent;

@Data
@ToString
public class OrderPlaced extends AbstractEvent {

    private Long id;
    private String testttttttttttttttttttttttt;
    private String name;
    private String price;
    private String qty;

    public OrderPlaced(Order aggregate) {
        super(aggregate);
    }

    public OrderPlaced() {
        super();
    }
}
