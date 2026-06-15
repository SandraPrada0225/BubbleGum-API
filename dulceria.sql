-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema dulceria
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema dulceria
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `dulceria` DEFAULT CHARACTER SET utf8 ;
-- -----------------------------------------------------
-- Schema new_schema1
-- -----------------------------------------------------
USE `dulceria` ;

-- -----------------------------------------------------
-- Table `dulceria`.`estados_carrito`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dulceria`.`estados_carrito` (
  `id` BIGINT NOT NULL,
  `nombre` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dulceria`.`carritos`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dulceria`.`carritos` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `subtotal` DECIMAL NOT NULL,
  `estado_carrito_id` BIGINT NOT NULL,
  `precio_total` DECIMAL NOT NULL,
  `envio` DECIMAL NOT NULL,,
  `descuento` DECIMAL NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_carritos_estados_carrito1_idx` (`estado_carrito_id` ASC) VISIBLE,
  CONSTRAINT `fk_carritos_estados_carrito1`
    FOREIGN KEY (`estado_carrito_id`)
    REFERENCES `dulceria`.`estados_carrito` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dulceria`.`categorias`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dulceria`.`categorias` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `nombre` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dulceria`.`usuarios`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dulceria`.`usuarios` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `nombre` VARCHAR(45) NULL,
  `apellido` VARCHAR(45) NULL,
  `password` VARCHAR(45) NULL,
  `correo` VARCHAR(45) NULL,
  `carrito_actual_id` INT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dulceria`.`presentaciones`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dulceria`.`presentaciones` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `nombre` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dulceria`.`marcas`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dulceria`.`marcas` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `nombre` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dulceria`.`dulces`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dulceria`.`dulces` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `presentacion_id` BIGINT NOT NULL,
  `marca_id` BIGINT NOT NULL,
  `nombre` VARCHAR(45) NOT NULL,
  `precio` DECIMAL NOT NULL,
  `peso` DECIMAL NOT NULL,
  `unidades` INT NOT NULL,
  `descripcion` VARCHAR(45) NOT NULL,
  `fecha_vencimiento` DATETIME NOT NULL,
  `fecha_expedicion` DATETIME NOT NULL,
  `disponibles` INT NOT NULL,
  `codigo` VARCHAR(45) NOT NULL,
  `imagen` VARCHAR(45) NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_dulces_presentaciones_idx` (`presentacion_id` ASC) VISIBLE,
  INDEX `fk_dulces_marcas1_idx` (`marca_id` ASC) VISIBLE,
  CONSTRAINT `fk_dulces_presentaciones`
    FOREIGN KEY (`presentacion_id`)
    REFERENCES `dulceria`.`presentaciones` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_dulces_marcas1`
    FOREIGN KEY (`marca_id`)
    REFERENCES `dulceria`.`marcas` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dulceria`.`medios_de_pago`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dulceria`.`medios_de_pago` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `nombre` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dulceria`.`ventas`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dulceria`.`ventas` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `medio_de_pago_id` BIGINT NOT NULL,
  `carrito_id` BIGINT NOT NULL,
  `comprador_id` BIGINT NOT NULL,
  `created_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_ventas_medios_de_pago1_idx` (`medio_de_pago_id` ASC) VISIBLE,
  INDEX `fk_ventas_carritos1_idx` (`carrito_id` ASC) VISIBLE,
  INDEX `fk_ventas_usuarios1_idx` (`comprador_id` ASC) VISIBLE,
  CONSTRAINT `fk_ventas_medios_de_pago1`
    FOREIGN KEY (`medio_de_pago_id`)
    REFERENCES `dulceria`.`medios_de_pago` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_ventas_carritos1`
    FOREIGN KEY (`carrito_id`)
    REFERENCES `dulceria`.`carritos` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_ventas_usuarios1`
    FOREIGN KEY (`comprador_id`)
    REFERENCES `dulceria`.`usuarios` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dulceria`.`categorias_dulces`
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
    REFERENCES `dulceria`.`dulces` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_categorias_dulces_categorias1`
    FOREIGN KEY (`categorias_id`)
    REFERENCES `dulceria`.`categorias` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dulceria`.`carritos_dulces`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dulceria`.`carritos_dulces` (
  `dulce_id` BIGINT NOT NULL,
  `carrito_id` BIGINT NOT NULL,
  `id` BIGINT NOT NULL,
  `unidades` INT NOT NULL,
  `subtotal` DECIMAL NOT NULL,
  INDEX `fk_carritos_dulces_dulces1_idx` (`dulce_id` ASC) VISIBLE,
  INDEX `fk_carritos_dulces_carritos1_idx` (`carrito_id` ASC) VISIBLE,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_carritos_dulces_dulces1`
    FOREIGN KEY (`dulce_id`)
    REFERENCES `dulceria`.`dulces` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_carritos_dulces_carritos1`
    FOREIGN KEY (`carrito_id`)
    REFERENCES `dulceria`.`carritos` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;


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
        ON d.marca_id = m.id
    INNER JOIN presentaciones p
        ON d.presentacion_id = p.id

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
        ON d.id = cd.dulce_id
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
        ON d.marca_id = m.id
    INNER JOIN presentaciones p
        ON d.presentacion_id = p.id

    WHERE d.id = p_id
    LIMIT 1;

END$$

DELIMITER ;
_____________

DROP PROCEDURE IF EXISTS GetPurchaseListByUserID;

DELIMITER $$

CREATE PROCEDURE GetPurchaseListByUserID(
    IN p_usuario_id BIGINT
)
BEGIN

    SELECT
        v.id,
        c.id AS carrito_id,
        v.created_at AS fecha,
        mp.id AS medio_de_pago_id,
        mp.nombre AS medio_de_pago,
        c.precio_total,
        c.subtotal,
        c.descuento,
        c.envio,
        ec.id AS estado_carrito_id,
        ec.nombre AS estado_carrito
    FROM ventas v
        INNER JOIN carritos c
            ON c.id = v.carrito_id
        INNER JOIN medios_de_pago mp
            ON mp.id = v.medio_de_pago_id
        INNER JOIN estados_carrito ec
            ON ec.id = c.estado_carrito_id
    WHERE v.comprador_id = p_usuario_id
    ORDER BY v.created_at DESC;

END$$

DELIMITER ;