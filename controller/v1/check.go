package v1

import (
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Reynadi531/phising-checker-api/utils"
	"github.com/gofiber/fiber/v2"
)

type CheckResult struct {
	IsFound      bool      `json:"isFound"`
	IsPhising    bool      `json:"isPhising"`
	IsSuspicious bool      `json:"isSuspicious"`
	Domain       string    `json:"domain"`
	Date         time.Time `json:"date"`
}

type ResponseCheck struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    CheckResult `json:"data"`
}

func CheckController(c *fiber.Ctx) error {
	if len(c.Query("url")) == 0 {
		return CheckControllerWithoutParam(c)
	}

	var cr CheckResult
	var rc ResponseCheck

	u, err := url.ParseRequestURI(c.Query("url"))
	if err != nil {
		cr = CheckResult{
			IsFound:   false,
			IsPhising: false,
			Domain:    "",
			Date:      time.Now(),
		}

		rc = ResponseCheck{
			Status:  http.StatusBadRequest,
			Message: "Failed to parse the url, please check your url",
			Data:    cr,
		}

		c.SendStatus(http.StatusBadRequest)
		return c.JSON(rc)
	}

	parts := strings.Split(u.Hostname(), ".")
	domain := parts[len(parts)-2] + "." + parts[len(parts)-1]

	phising, err := utils.CheckLink(domain)
	if err != nil {
		panic(err)
	}

	sus, err := utils.CheckSusDomain(domain)
	if err != nil {
		panic(err)
	}

	if sus && phising {
		cr = CheckResult{
			IsFound:      true,
			IsPhising:    true,
			IsSuspicious: true,
			Domain:       domain,
			Date:         time.Now(),
		}

		rc = ResponseCheck{
			Status:  http.StatusOK,
			Message: "Found, its suspicious and phising",
			Data:    cr,
		}

		c.Status(http.StatusOK)
		return c.JSON(rc)
	} else if sus {
		cr = CheckResult{
			IsFound:      true,
			IsPhising:    false,
			IsSuspicious: true,
			Domain:       domain,
			Date:         time.Now(),
		}

		rc = ResponseCheck{
			Status:  http.StatusOK,
			Message: "Found, its suspicious",
			Data:    cr,
		}

		c.Status(http.StatusOK)
		return c.JSON(rc)
	} else if phising {
		cr = CheckResult{
			IsFound:      true,
			IsPhising:    true,
			IsSuspicious: false,
			Domain:       domain,
			Date:         time.Now(),
		}

		rc = ResponseCheck{
			Status:  http.StatusOK,
			Message: "Found, its phising",
			Data:    cr,
		}

		c.Status(http.StatusOK)
		return c.JSON(rc)
	}
	return nil
}

func CheckControllerWithoutParam(c *fiber.Ctx) error {

	cr := CheckResult{
		IsFound:      false,
		IsPhising:    false,
		IsSuspicious: false,
		Domain:       "",
		Date:         time.Now(),
	}

	rc := ResponseCheck{
		Status:  http.StatusBadRequest,
		Message: "Parameter not inclueded",
		Data:    cr,
	}

	c.SendStatus(http.StatusBadRequest)
	return c.JSON(rc)
}
