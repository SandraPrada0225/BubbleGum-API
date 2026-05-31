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
  `id` BIGINT NOT NULL,
  `estados_carrito_id` BIGINT NOT NULL,
  `precio_total` DECIMAL NOT NULL,
  `fecha` DATE NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_carritos_estados_carrito1_idx` (`estados_carrito_id` ASC) VISIBLE,
  CONSTRAINT `fk_carritos_estados_carrito1`
    FOREIGN KEY (`estados_carrito_id`)
    REFERENCES `dulceria`.`estados_carrito` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dulceria`.`categorias`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dulceria`.`categorias` (
  `id` BIGINT NOT NULL,
  `nombre` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dulceria`.`usuarios`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dulceria`.`usuarios` (
  `id` BIGINT NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dulceria`.`presentaciones`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dulceria`.`presentaciones` (
  `id` BIGINT NOT NULL,
  `nombre` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dulceria`.`marcas`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dulceria`.`marcas` (
  `id` BIGINT NOT NULL,
  `nombre` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dulceria`.`dulces`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dulceria`.`dulces` (
  `id` BIGINT NOT NULL,
  `presentaciones_id` BIGINT NOT NULL,
  `marcas_id` BIGINT NOT NULL,
  `nombre` VARCHAR(45) NOT NULL,
  `precio` DECIMAL NOT NULL,
  `peso` DECIMAL NOT NULL,
  `unidades` INT NOT NULL,
  `descripcion` LONGTEXT NOT NULL,
  `fecha_vencimiento` DATETIME NOT NULL,
  `fecha_expedicion` DATETIME NOT NULL,
  `disponibles` INT NOT NULL,
  `codigo` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_dulces_presentaciones_idx` (`presentaciones_id` ASC) VISIBLE,
  INDEX `fk_dulces_marcas1_idx` (`marcas_id` ASC) VISIBLE,
  CONSTRAINT `fk_dulces_presentaciones`
    FOREIGN KEY (`presentaciones_id`)
    REFERENCES `dulceria`.`presentaciones` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_dulces_marcas1`
    FOREIGN KEY (`marcas_id`)
    REFERENCES `dulceria`.`marcas` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dulceria`.`medios_de_pago`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dulceria`.`medios_de_pago` (
  `id` BIGINT NOT NULL,
  `nombre` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dulceria`.`ventas`
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
    REFERENCES `dulceria`.`medios_de_pago` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_ventas_carritos1`
    FOREIGN KEY (`carritos_id`)
    REFERENCES `dulceria`.`carritos` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_ventas_usuarios1`
    FOREIGN KEY (`usuarios_id`)
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
    REFERENCES `dulceria`.`dulces` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_carritos_dulces_carritos1`
    FOREIGN KEY (`carritos_id`)
    REFERENCES `dulceria`.`carritos` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
