-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `dulceria` DEFAULT CHARACTER SET utf8 ;
-- -----------------------------------------------------
-- Schema new_schema1
-- -----------------------------------------------------
USE `dulceria` ;

-- -----------------------------------------------------
-- Table `mydb`.`estados_carrito`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dulceria`.`estados_carrito` (
  `id` BIGINT NOT NULL,
  `nombre` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`carritos`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dulceria`.`carritos` (
  `id` BIGINT NOT NULL,
  `estados_carrito_id` BIGINT NOT NULL,
  `precio_total` DECIMAL NOT NULL,
  `fecha` DATE NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_carritos_estados_carrito1_idx` (`estados_carrito_id` ASC) VISIBLE,
  CONSTRAINT `fk_carritos_estados_carrito1`
    FOREIGN KEY (`estados_carrito_id`)
    REFERENCES `mydb`.`estados_carrito` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`categorias`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dulceria`.`categorias` (
  `id` BIGINT NOT NULL,
  `nombre` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`usuarios`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dulceria`.`usuarios` (
  `id` BIGINT NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`presentaciones`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dulceria`.`presentaciones` (
  `id` BIGINT NOT NULL,
  `nombre` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`marcas`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dulceria`.`marcas` (
  `id` BIGINT NOT NULL,
  `nombre` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`dulces`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dulceria`.`dulces` (
  `id` BIGINT NOT NULL,
  `presentaciones_id` BIGINT NOT NULL,
  `marcas_id` BIGINT NOT NULL,
  `nombre` VARCHAR(45) NOT NULL,
  `precio` DECIMAL NOT NULL,
  `peso` DECIMAL NOT NULL,
  `unidades` INT NOT NULL,
  `descripcion` VARCHAR(45) NOT NULL,
  `fecha_vencimiento` DATETIME NOT NULL,
  `fecha_expedicion` DATETIME NOT NULL,
  `disponibles` INT NOT NULL,
  `codigo` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_dulces_presentaciones_idx` (`presentaciones_id` ASC) VISIBLE,
  INDEX `fk_dulces_marcas1_idx` (`marcas_id` ASC) VISIBLE,
  CONSTRAINT `fk_dulces_presentaciones`
    FOREIGN KEY (`presentaciones_id`)
    REFERENCES `mydb`.`presentaciones` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_dulces_marcas1`
    FOREIGN KEY (`marcas_id`)
    REFERENCES `mydb`.`marcas` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`medios_de_pago`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dulceria`.`medios_de_pago` (
  `id` BIGINT NOT NULL,
  `nombre` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`ventas`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dulceria`.`ventas` (
  `id` BIGINT NOT NULL,
  `medios_de_pago_id` BIGINT NOT NULL,
  `carritos_id` BIGINT NOT NULL,
  `usuarios_id` BIGINT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_ventas_medios_de_pago1_idx` (`medios_de_pago_id` ASC) VISIBLE,
  INDEX `fk_ventas_carritos1_idx` (`carritos_id` ASC) VISIBLE,
  INDEX `fk_ventas_usuarios1_idx` (`usuarios_id` ASC) VISIBLE,
  CONSTRAINT `fk_ventas_medios_de_pago1`
    FOREIGN KEY (`medios_de_pago_id`)
    REFERENCES `mydb`.`medios_de_pago` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_ventas_carritos1`
    FOREIGN KEY (`carritos_id`)
    REFERENCES `mydb`.`carritos` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_ventas_usuarios1`
    FOREIGN KEY (`usuarios_id`)
    REFERENCES `mydb`.`usuarios` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`categorias_dulces`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dulceria`.`categorias_dulces` (
  `dulces_id` BIGINT NOT NULL,
  `categorias_id` BIGINT NOT NULL,
  `id` BIGINT NOT NULL,
  INDEX `fk_categorias_dulces_dulces1_idx` (`dulces_id` ASC) VISIBLE,
  INDEX `fk_categorias_dulces_categorias1_idx` (`categorias_id` ASC) VISIBLE,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_categorias_dulces_dulces1`
    FOREIGN KEY (`dulces_id`)
    REFERENCES `mydb`.`dulces` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_categorias_dulces_categorias1`
    FOREIGN KEY (`categorias_id`)
    REFERENCES `mydb`.`categorias` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`carritos_dulces`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dulceria`.`carritos_dulces` (
  `dulces_id` BIGINT NOT NULL,
  `carritos_id` BIGINT NOT NULL,
  `id` BIGINT NOT NULL,
  `unidades` INT NOT NULL,
  `subtotal` DECIMAL NOT NULL,
  INDEX `fk_carritos_dulces_dulces1_idx` (`dulces_id` ASC) VISIBLE,
  INDEX `fk_carritos_dulces_carritos1_idx` (`carritos_id` ASC) VISIBLE,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_carritos_dulces_dulces1`
    FOREIGN KEY (`dulces_id`)
    REFERENCES `mydb`.`dulces` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_carritos_dulces_carritos1`
    FOREIGN KEY (`carritos_id`)
    REFERENCES `mydb`.`carritos` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;

ALTER TABLE dulces
ADD COLUMN imagen VARCHAR(255) NOT NULL,
ADD COLUMN subtotal INT NOT NULL DEFAULT 0;

-- -----------------------------------------------------
-- procedimiento almacenado GetDetalleDulceByCode
-- -----------------------------------------------------
DELIMITER $$

CREATE PROCEDURE GetDetalleDulceByCode(
    IN pCodigo VARCHAR(45)
)
BEGIN

    SELECT
        d.id,
        d.peso,
        d.precio AS precio_unidad,
        d.disponibles,
        0 AS subtotal,
        d.codigo,
        d.nombre,
        d.descripcion,
        d.imagen,
        DATE_FORMAT(d.fecha_vencimiento, '%Y-%m-%d') AS fecha_vencimiento,
        DATE_FORMAT(d.fecha_expedicion, '%Y-%m-%d') AS fecha_expedicion,

        m.id AS marca_id,
        m.nombre AS marca_nombre,

        p.id AS presentacion_id,
        p.nombre AS presentacion_nombre

    FROM dulces d
    INNER JOIN marcas m
        ON d.marcas_id = m.id
    INNER JOIN presentaciones p
        ON d.presentaciones_id = p.id

    WHERE d.codigo = pCodigo
    LIMIT 1;

END$$

DELIMITER ;

-- -----------------------------------------------------
-- procedimiento almacenado GetCategoriasDulceCode
-- -----------------------------------------------------

DELIMITER $$

CREATE PROCEDURE GetCategoriasByDulceCode(
    IN p_codigo VARCHAR(45)
)
BEGIN

    SELECT
        c.id,
        c.nombre
    FROM categorias c
    INNER JOIN categorias_dulces cd
        ON cd.categorias_id = c.id
    INNER JOIN dulces d
        ON d.id = cd.dulces_id
    WHERE d.codigo = p_codigo;

END$$

DELIMITER ;


DELIMITER $$

CREATE PROCEDURE GetCategoriasByDulceID(
    IN p_id VARCHAR(45)
)
BEGIN

    SELECT
        c.id,
        c.nombre
    FROM categorias c
    INNER JOIN categorias_dulces cd
        ON cd.categorias_id = c.id
    INNER JOIN dulces d
        ON d.id = cd.dulces_id
    WHERE d.id = p_id;

END$$

DELIMITER ;

-- -----------------------------------------------------
-- procedimiento almacenado GetDetalleDulceByID
-- -----------------------------------------------------
DELIMITER $$

CREATE PROCEDURE GetDetalleDulceByID(
    IN p_id int 
)
BEGIN

    SELECT
        d.id,
        d.peso,
        d.precio AS precio_unidad,
        d.disponibles,
        0 AS subtotal,
        d.codigo,
        d.nombre,
        d.descripcion,
        d.imagen,
        DATE_FORMAT(d.fecha_vencimiento, '%Y-%m-%d') AS fecha_vencimiento,
        DATE_FORMAT(d.fecha_expedicion, '%Y-%m-%d') AS fecha_expedicion,

        m.id AS marca_id,
        m.nombre AS marca_nombre,

        p.id AS presentacion_id,
        p.nombre AS presentacion_nombre

    FROM dulces d
    INNER JOIN marcas m
        ON d.marcas_id = m.id
    INNER JOIN presentaciones p
        ON d.presentaciones_id = p.id

    WHERE d.id = p_id
    LIMIT 1;

END$$

DELIMITER ;
_____________
--agregamos los campos subtotal y descuento al carrito
ALTER TABLE carritos
ADD COLUMN subtotal DECIMAL(10,2) NOT NULL DEFAULT 0,
ADD COLUMN descuento DECIMAL(10,2) NOT NULL DEFAULT 0;


ALTER TABLE usuarios
ADD COLUMN nombre VARCHAR(100) NOT NULL,
ADD COLUMN apellido VARCHAR(100) NOT NULL,
ADD COLUMN correo VARCHAR(150) UNIQUE,
ADD COLUMN telefono VARCHAR(20);


ALTER TABLE ventas
ADD COLUMN created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE carritos
ADD COLUMN envio DECIMAL(10,2) NOT NULL DEFAULT 0;

ALTER TABLE ventas
MODIFY COLUMN id BIGINT NOT NULL AUTO_INCREMENT;