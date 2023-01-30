package testing.domain;

import java.util.Date;
import java.util.List;
import javax.persistence.*;
import lombok.Data;
import testing.TestApplication;
import testing.domain.OrderPlaced;

@Entity
@Table(name = "Order_table")
@Data
public class Order {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Long id;

    private String testttttttttttttttttttttttt;

    private String name;

    private String price;

    private String qty;

    @PostPersist
    public void onPostPersist() {
        OrderPlaced orderPlaced = new OrderPlaced(this);
        orderPlaced.publishAfterCommit();
    }

    public static OrderRepository repository() {
        OrderRepository orderRepository = TestApplication.applicationContext.getBean(
            OrderRepository.class
        );
        return orderRepository;
    }
}
